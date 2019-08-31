package approvalsvc

import (
	"context"
	"stash-mono-repo/service/approvalsvc/approval_logic"
	"stash-mono-repo/service/approvalsvc/model"
)

// Service provides some "date capabilities" to your application
type Service interface {
	Status(ctx context.Context) (string, error)
	GetApprovals(ctx context.Context, req model.GetApprovalsRequest) (model.GetApprovalsResponse, error)
	AddApproval(ctx context.Context, req model.AddApprovalRequest) (model.AddApprovalResponse, error)
	UpdateApproval(ctx context.Context, req model.UpdateApprovalRequest) (model.UpdateApprovalResponse, error)
}

type approvalService struct{}

// NewService makes a new Service.
func NewService() Service {
	return approvalService{}
}

// Status only tell us that our service is ok!
func (approvalService) Status(ctx context.Context) (string, error) {
	return "ok", nil
}

// Get will return all approvals
func (approvalService) GetApprovals(ctx context.Context, req model.GetApprovalsRequest) (resp model.GetApprovalsResponse, err error) {
	return approval_logic.GetApprovals(ctx, req)
}

// Add will add a new approval
func (approvalService) AddApproval(ctx context.Context, req model.AddApprovalRequest) (resp model.AddApprovalResponse, err error) {
	return approval_logic.AddApproval(ctx, req)
}

// Update will update an existing approval
func (approvalService) UpdateApproval(ctx context.Context, req model.UpdateApprovalRequest) (resp model.UpdateApprovalResponse, err error) {
	return approval_logic.UpdateApproval(ctx, req)
}
