package approval_logic

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"stash-mono-repo/service/approvalsvc/model"

	"github.com/jinzhu/gorm"
)

/*
	Steps
	1) Check request and make sure that everything is valid (Service Rule validity is not implemented programmatically but in database FK check)
	2) Retreive current approval object from database and check to see if ID exists
	3) Do a status check so that only comments can be updated for approved / rejected and so on...
	4) [Conditional] If the update is for status, we call the designated service (service rule)
	5) Update database accordingly
	6) Send updated result back to client
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
	}

	// ------------ STEP 4 --------------
	/*
		Approval Status has changed from pending / error to approved / rejected / cancelled! We need to inform the designated service!
		Now we send the update to the defined service rule.
		All request will be sent via the same format to any other approval receipient.
	*/
	var serverErr error
	if hasApprovalStatusChanged {
		serviceRuleID := approvalItem.ServiceRule
		if req.ServiceRule != nil {
			serviceRuleID = *req.ServiceRule
		}

		serviceRule, err := getServiceRule(db, serviceRuleID)
		fmt.Println(err)

		approvalToSend := model.ApprovalToSend{
			ID:     id,
			Status: *req.Status,
		}

		serverErr = sendRequestToServiceRule(serviceRule, approvalToSend)
	}

	// ------------ STEP 5 --------------
	if serverErr != nil {
		toUpdate["status"] = model.STATUS_ERROR
		if toUpdate["comment"] != nil {
		}

		//toUpdate["comment"]

	}
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
		toUpdate["serviceRule"] = req.ServiceRule
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
			err = errors.New(model.PQ_ERROR_NO_ROWS_FOUND)
			return
		}

		fmt.Println(tmpDB.Error) //TODO: Move to log service
		err = errors.New(model.ERROR_DATABASE_ERROR)
		return
	}
	return
}

// HTTP Call to the required service once all checks have been passed and approval has changed!
func sendRequestToServiceRule(serviceRule model.ServiceRule, approvalToSend model.ApprovalToSend) (err error) {
	url := "http://restapi3.apiary.io/notes"
	fmt.Println("URL:>", url)

	var jsonStr = []byte(`{"title":"Buy cheese and bread for breakfast."}`)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
	return
}

// We update the table here
func updateApprovalTable(db *gorm.DB, toUpdate map[string]interface{}, id string) (updatedModel model.ApprovalItem, err error) {
	db.Debug()
	db.LogMode(true)
	//a := model.ApprovalItem{}
	tmpDB := db.Table(model.APPROVAL_TABLE).Where("id = ?", id).Updates(toUpdate)
	if tmpDB.Error != nil {
		fmt.Println(tmpDB.Error) //TODO: Move to log service
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

/*
// Copy latest variables to a new approval object to reply back to client
func updateNewApprovalItem(approvalItem model.ApprovalItem, toUpdate map[string]interface{}) (newApprovalItem model.ApprovalItem) {
	newApprovalItem.ID = approvalItem.ID
	if val, ok := toUpdate["status"]; ok {
		newApprovalItem.Status = *(val.(*int))
	} else {
		newApprovalItem.Status = approvalItem.Status
	}

	if val, ok := toUpdate["title"]; ok {
		newApprovalItem.Title = *(val.(*string))
	} else {
		newApprovalItem.Title = approvalItem.Title
	}

	if val, ok := toUpdate["description"]; ok {
		newApprovalItem.Description = *(val.(*string))
	} else {
		newApprovalItem.Description = approvalItem.Description
	}

	if val, ok := toUpdate["serviceRule"]; ok {
		newApprovalItem.ServiceRule = *(val.(*int))
	} else {
		newApprovalItem.ServiceRule = approvalItem.ServiceRule
	}

	if val, ok := toUpdate["deadline"]; ok {
		newApprovalItem.Deadline = val.(*time.Time)
	} else {
		newApprovalItem.Deadline = approvalItem.Deadline
	}

	if val, ok := toUpdate["comment"]; ok {
		newApprovalItem.Comment = *(val.(*string))
	} else {
		newApprovalItem.Comment = approvalItem.Comment
	}

	return
}
*/
