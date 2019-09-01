package approval_logic

import (
	"context"
	"errors"
	"fmt"
	"log"
	"stash-mono-repo/service/approvalsvc/model"

	"github.com/jinzhu/gorm"
)

func GetApprovals(ctx context.Context, req model.GetApprovalsRequest, db *gorm.DB, logger log.Logger) (resp model.GetApprovalsResponse, err error) {
	approvalItems := []model.ApprovalItem{}
	tmpDB := db

	// TODO: Improve this into a proper sorting / where query.
	switch req.Default {
	case 0:
		tmpDB = db.Table(model.APPROVAL_TABLE).Order("deadline ASC").Where("status = ?", model.STATUS_PENDING)
	case 1:
		tmpDB = db.Table(model.APPROVAL_TABLE).Order("deadline ASC")
	}

	if req.Offset != -1 {
		tmpDB = tmpDB.Offset(req.Offset)
	}

	if req.Limit != -1 {
		tmpDB = tmpDB.Limit(req.Limit)
	}

	tmpDB.Find(&approvalItems)

	if tmpDB.Error != nil {
		fmt.Println(tmpDB.Error) //TODO: Move to log service
		err = errors.New(model.ERROR_DATABASE_ERROR)
		return
	}
	resp = model.GetApprovalsResponse{
		ApprovalItems: approvalItems,
	}

	return
}
