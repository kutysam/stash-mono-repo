package approval_logic

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"stash-mono-repo/service/approvalsvc/model"
)

const (
	sqlSelectStatement = `SELECT * FROM approval`
)

func GetApprovals(ctx context.Context, req model.GetApprovalsRequest, db *sql.DB, logger log.Logger) (resp model.GetApprovalsResponse, err error) {
	var rows *sql.Rows
	rows, err = db.Query(sqlSelectStatement)
	approvalItems := []model.ApprovalItem{}

	for rows.Next() {
		tmp := model.ApprovalItem{}
		err = rows.Scan(&tmp.ID, &tmp.Title, &tmp.Description, &tmp.ServiceRule, &tmp.Comment, &tmp.Status, &tmp.Deadline, &tmp.CreatedAt, &tmp.UpdatedAt)
		if err != nil {
			fmt.Println(err) //Change to log service when available
			err = errors.New(model.ERROR_DATABASE_ERROR)
		}
		approvalItems = append(approvalItems, tmp)
	}

	resp = model.GetApprovalsResponse{
		ApprovalItems: approvalItems,
	}

	return
}
