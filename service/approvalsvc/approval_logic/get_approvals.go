package approval_logic

import (
	"context"
	"stash-mono-repo/service/approvalsvc/model"
)

func GetApprovals(ctx context.Context, req model.GetApprovalsRequest) (resp model.GetApprovalsResponse, err error) {
	approvalItems := []model.ApprovalItem{}
	approvalItem := model.ApprovalItem{
		ID:          "aa",
		ServiceRule: 4,
		Priority:    2,
		Comment:     "Comment",
		Status:      "Pending",
	}
	approvalItems = append(approvalItems, approvalItem)

	resp = model.GetApprovalsResponse{
		ApprovalItems: approvalItems,
		Err:           "no error",
	}

	return
}
