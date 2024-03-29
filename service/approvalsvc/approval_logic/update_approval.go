package approval_logic

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"stash-mono-repo/service/approvalsvc/model"

	"github.com/jinzhu/gorm"
	"github.com/lib/pq"
)

/*
	Steps
	1) Check request and make sure that everything is valid (Service Rule validity is not implemented programmatically but in database FK check)
	2) Retreive current approval object from database and check to see if ID exists
	3) Do a status check so that only comments can be updated for approved / rejected and so on...
	4) [Conditional] If the update is for status, we call the designated service (service rule)
	5) Update database accordingly [If it is an update for status, we will mark it as acknowledged_approved / acknowledged_rejected / acknowledged_cancelled]
	6) Send updated result back to client
	7) We will call a goroutine (New thread) to update the designated service if (4 is true)
	8) Update databse accordingly after this.
*/

func UpdateApproval(ctx context.Context, req model.UpdateApprovalRequest, db *gorm.DB, logger log.Logger) (resp model.UpdateApprovalResponse, err error) {
	// ------------ STEP 1 --------------
	// Check request to see if it is valid
	id := req.ID
	if id == "" {
		err = errors.New(model.ERROR_NO_ID_PROVIDED)
		return
	}

	// We update only what we need to update.
	var toUpdate map[string]interface{}
	toUpdate, err = fillInUpdateFields(req)
	if err != nil {
		return
	}

	if len(toUpdate) == 0 {
		err = errors.New(model.ERROR_NO_FIELDS_TO_UPDATE)
		return
	}

	// ------------ STEP 2 --------------
	// Now we get the same ID from the database.
	approvalItem := model.ApprovalItem{}
	tmpDB := db.Table(model.APPROVAL_TABLE).Where("id = ?", req.ID).First(&approvalItem)

	if tmpDB.Error != nil {
		if tmpDB.Error == gorm.ErrRecordNotFound {
			err = errors.New(model.PQ_ERROR_NO_ROWS_FOUND)
			return
		}

		fmt.Println(tmpDB.Error) //TODO: Move to log service
		err = errors.New(model.ERROR_DATABASE_ERROR)
		return
	}

	// The following is solely done for the server side to receive the approval item with the updated fields (If all goes well)
	var updatedApprovalItem model.ApprovalItem

	// ------------ STEP 3 --------------
	// Check approval status accordingly
	hasApprovalStatusChanged := false
	// If the item has been approved or rejected, you can update the comment and nothing else!
	// If the user sends a request with other fields change, the request will be rejected. You can only send id, comment and nothing else!
	if approvalItem.Status == model.STATUS_APPROVED || approvalItem.Status == model.STATUS_REJECTED || approvalItem.Status == model.STATUS_CANCELLED {
		if len(toUpdate) == 1 && req.Comment != nil {
			// We update the comment here
			updatedApprovalItem, err = updateApprovalTable(db, toUpdate, approvalItem.ID)
			if err != nil {
				return
			} else {
				resp.ApprovalItem = updatedApprovalItem
				return
			}
		} else {
			err = errors.New(model.ERROR_COMMENT_UPDATE_ONLY) //If the user placed in other fields in the update request, it will not go through and user will have to resubmit the request.
			return
		}
	} else if approvalItem.Status == model.STATUS_PENDING || approvalItem.Status == model.STATUS_ERROR {
		if req.Status != nil && approvalItem.Status != *req.Status {
			hasApprovalStatusChanged = true
			var tmpStatus int
			if *req.Status == model.STATUS_ACKNOWLEDGED_APPROVED {
				tmpStatus = model.STATUS_ACKNOWLEDGED_APPROVED
			} else if *req.Status == model.STATUS_REJECTED {
				tmpStatus = model.STATUS_ACKNOWLEDGED_REJECTED
			} else if *req.Status == model.STATUS_CANCELLED {
				tmpStatus = model.STATUS_ACKNOWLEDGED_CANCELLED
			} else {
				err = errors.New("You can only transition from error / pending --> approved / cancelled / rejected.")
				return
			}
			toUpdate["status"] = &tmpStatus
			updatedApprovalItem, err = updateApprovalTable(db, toUpdate, approvalItem.ID)
		} else {
			// We update the table here and return to the client as the status hasn't been changed!
			updatedApprovalItem, err = updateApprovalTable(db, toUpdate, approvalItem.ID)
			if err != nil {
				return
			} else {
				resp.ApprovalItem = updatedApprovalItem
				return
			}
		}
	} else {
		err = errors.New("Sorry, you cannot update the request. Your approval object in a state where we are waiting for the other service's response.")
		return
	}

	// ------------ STEP 4 --------------
	/*
		Approval Status has changed from pending / error to approved / rejected / cancelled! We need to inform the designated service!
		Now we send the update to the defined service rule.
		All request will be sent via the same format to any other approval receipient.
	*/

	if hasApprovalStatusChanged {
		serviceRuleID := approvalItem.ServiceRule
		if req.ServiceRule != nil {
			serviceRuleID = *req.ServiceRule
		}

		var serviceRule model.ServiceRule
		serviceRule, err = getServiceRule(db, serviceRuleID)
		if err != nil {
			return
		}

		approvalToSend := model.SendChangedApprovalRequest{
			ID:     id,
			Status: *req.Status,
		}

		// Call a go routine to thread branch out a thread for the httpcall.
		// Meanwhile, we will close the connection with the approver by sending him a acknowledged approved / rejected / cancelled status
		go sendRequestToServiceRule(db, serviceRule, approvalToSend, toUpdate)
	}

	resp.ApprovalItem = updatedApprovalItem
	return
}

func fillInUpdateFields(req model.UpdateApprovalRequest) (map[string]interface{}, error) {
	toUpdate := make(map[string]interface{})
	if req.Status != nil {
		if model.CheckValidStatus(*req.Status) {
			toUpdate["status"] = req.Status
		} else {
			return toUpdate, errors.New(model.ERROR_INVALID_STATUS)
		}
	}
	if req.Title != nil {
		toUpdate["title"] = req.Title
	}
	if req.Description != nil {
		toUpdate["description"] = req.Description
	}
	if req.ServiceRule != nil {
		toUpdate["service_rule"] = req.ServiceRule
	}
	if req.Deadline != nil {
		toUpdate["deadline"] = req.Deadline
	}
	if req.Comment != nil {
		toUpdate["comment"] = req.Comment
	}
	return toUpdate, nil
}

// Retreive the service rule objects accordingly
func getServiceRule(db *gorm.DB, serviceRuleID int) (serviceRule model.ServiceRule, err error) {
	tmpDB := db.Table(model.SERVICE_RULE_TABLE).Where("id = ?", serviceRuleID).First(&serviceRule)
	if tmpDB.Error != nil {
		if tmpDB.Error == gorm.ErrRecordNotFound {
			err = errors.New(model.PQ_SUCH_SERVICE_RULE_FOUND)
			return
		}

		fmt.Println(tmpDB.Error) //TODO: Move to log service
		err = errors.New(model.ERROR_DATABASE_ERROR)
		return
	}
	return
}

// HTTP Call to the required service once all checks have been passed and approval has changed!
func sendRequestToServiceRule(db *gorm.DB, serviceRule model.ServiceRule, approvalToSend model.SendChangedApprovalRequest, toUpdate map[string]interface{}) (err error) {
	url := serviceRule.URL

	approvalToSendBytes, err := json.Marshal(approvalToSend)
	if err != nil {
		fmt.Println(err) // TODO: Move to log service
		err = errors.New(model.ERROR_UNABLE_TO_MARSHAL)
		updateErrorToDB(db, toUpdate, approvalToSend, err)
		return
	}

	request, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(approvalToSendBytes))
	if err != nil {
		fmt.Println(err)
		err = errors.New(model.ERROR_UPDATING_SERVICE)
		updateErrorToDB(db, toUpdate, approvalToSend, err)
		return
	}
	request.Header.Add("Authorization", serviceRule.Apikey)
	request.Header.Add("Content-type", "application/json")

	client := &http.Client{}

	resp, err := client.Do(request) // TODO: Do a maximum of 3 retries if it fails.
	if err != nil {
		fmt.Println(err) // TODO: Move to log service
		err = errors.New(model.ERROR_UPDATING_SERVICE)
		updateErrorToDB(db, toUpdate, approvalToSend, err)
		return
	}

	if resp.StatusCode != 200 {
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Print(resp.StatusCode) // TODO: Move to log service
		fmt.Println(string(body))  // TODO: Move to log service
		err = errors.New(model.ERROR_UPDATING_SERVICE)
		updateErrorToDB(db, toUpdate, approvalToSend, err)
		return
	} else {
		var result model.SendChangedApprovalResponse
		json.NewDecoder(resp.Body).Decode(&result)

		if result.Status != approvalToSend.Status {
			fmt.Println("Seems like the server has already received the approval status before and it is different from what you are sending. We will be using " + string(result.Status) + " approval status instead of your requested approval status: " + string(approvalToSend.Status) + " ERROR CODE: " + string(result.ErrorCode)) //TODO: Log this down
			toUpdate["status"] = result.Status
		} else {
			toUpdate["status"] = result.Status
		}
	}

	// ------------ STEP 5 --------------
	// We check if there is an error, if there is an error talking to the designated service, we stop and mark the request as errored and update the database.
	// The developer can do some slack notification etc. over here to immediately alert the people.
	// If there is no error, we just send a 200 OK and the updated fields back to client.
	updateErrorToDB(db, toUpdate, approvalToSend, err)

	return
}

func updateErrorToDB(db *gorm.DB, toUpdate map[string]interface{}, approvalToSend model.SendChangedApprovalRequest, err error) {
	if err != nil {
		fmt.Println(err) // TODO: Log this error down
		toUpdate["status"] = model.STATUS_ERROR
		toUpdate["comment"] = model.ERROR_UPDATING_SERVICE
		var err1 error
		_, err1 = updateApprovalTable(db, toUpdate, approvalToSend.ID)
		if err1 != nil {
			err1 = errors.New(model.ERROR_UPDATING_SERVICE)
			err = err1
		}
		return
	} else {
		// This is the best success that we achieved. No errors all the way. We update database with the final status (Either approved / rejected / cancelled)
		_, err = updateApprovalTable(db, toUpdate, approvalToSend.ID)
		if err != nil {
			err = errors.New(model.ERROR_UPDATING_SERVICE)
			return
		}
	}
}

// We update the table here
func updateApprovalTable(db *gorm.DB, toUpdate map[string]interface{}, id string) (updatedModel model.ApprovalItem, err error) {
	tmpDB := db.Table(model.APPROVAL_TABLE).Where("id = ?", id).Updates(toUpdate)
	if tmpDB.Error != nil {
		fmt.Println(tmpDB.Error) //TODO: Move to log service
		pqErr, ok := tmpDB.Error.(*pq.Error)
		if ok {
			if pqErr.Code == model.PQ_ERROR_FOREIGN_KEY_VIOLATION && pqErr.Constraint == model.PQ_SERVICE_RULE_CONSTRAINT {
				err = errors.New(model.ERROR_SERVICE_RULE_INVALID_CONSTRAINT)
				return
			}
		}
		err = errors.New(model.ERROR_DATABASE_ERROR)
		return
	}

	// Get the item from the last update and send it back
	tmpDB = tmpDB.First(&updatedModel)
	if tmpDB.Error != nil {
		fmt.Println(tmpDB.Error) //TODO: Move to log service
		err = errors.New(model.ERROR_DATABASE_ERROR)
		return
	}

	return
}
