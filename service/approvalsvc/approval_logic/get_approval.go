package approval_logic

import (
	"context"
	"errors"
	"fmt"
	"log"
	"stash-mono-repo/service/approvalsvc/model"

	"github.com/jinzhu/gorm"
)

func GetApproval(ctx context.Context, req model.GetApprovalRequest, db *gorm.DB, logger log.Logger) (resp model.GetApprovalResponse, err error) {
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

	resp = model.GetApprovalResponse{
		ApprovalItem: approvalItem,
	}
	return
}
