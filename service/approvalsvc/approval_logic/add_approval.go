package approval_logic

import (
	"context"
	"stash-mono-repo/service/approvalsvc/model"
)

func AddApproval(ctx context.Context, req model.AddApprovalRequest) (resp model.AddApprovalResponse, err error) {
	approvalItem := model.ApprovalItem{
		ID:          "azza",
		ServiceRule: 4,
		Priority:    2,
		Comment:     "Comment",
		Status:      "Pending",
	}

	resp = model.AddApprovalResponse{
		ApprovalItem: approvalItem,
		Err:          "no error",
	}

	return
}
