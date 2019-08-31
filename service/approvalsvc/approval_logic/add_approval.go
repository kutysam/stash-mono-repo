package approval_logic

import (
	"context"
	"errors"
	"stash-mono-repo/service/approvalsvc/model"
	"time"

	"github.com/google/uuid"
)

func AddApproval(ctx context.Context, req model.AddApprovalRequest) (resp model.AddApprovalResponse, err error) {
	priority := req.Priority
	serviceRule := req.ServiceRule
	deadline := &req.Deadline
	comment := req.Comment
	serviceRule = req.ServiceRule

	if _, ok := model.ServiceRule[serviceRule]; !ok {
		err = errors.New(model.ERROR_INVALID_SERVICE_RULE)
		return
	}

	if time.Now().After(*deadline) {
		err = errors.New(model.ERROR_CURRENT_TIME_IS_AFTER_DEADLINE)
		return
	}

	approvalItem := model.ApprovalItem{
		ID:          uuid.New().String(),
		ServiceRule: serviceRule,
		Priority:    priority,
		Comment:     comment,
		Status:      model.STATUS_PENDING,
		Deadline:    deadline,
	}

	resp = model.AddApprovalResponse{
		ApprovalItem: approvalItem,
		Err:          "nil",
	}

	return
}
