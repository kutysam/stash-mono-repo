package approval_logic

import (
	"context"
	"database/sql"
	"log"
	"stash-mono-repo/service/approvalsvc/model"
)

func UpdateApproval(ctx context.Context, req model.UpdateApprovalRequest, db *sql.DB, logger log.Logger) (resp model.UpdateApprovalResponse, err error) {
	approvalItem := model.ApprovalItem{
		ID:          "aa",
		ServiceRule: 4,
		//Priority:    2,
		Comment: "Comment",
		Status:  model.STATUS_PENDING,
	}

	resp = model.UpdateApprovalResponse{
		ApprovalItem: approvalItem,
		Err:          "no error",
	}

	return
}
