package approvalsvc

import (
	"context"
	"database/sql"
	"log"
	"stash-mono-repo/service/approvalsvc/approval_logic"
	"stash-mono-repo/service/approvalsvc/model"
)

// Service provides some "date capabilities" to your application
type Service interface {
	Status(ctx context.Context) (string, error)
	GetApprovals(ctx context.Context, req model.GetApprovalsRequest) (model.GetApprovalsResponse, error)
	AddApproval(ctx context.Context, req model.AddApprovalRequest, db *sql.DB) (model.AddApprovalResponse, error)
	UpdateApproval(ctx context.Context, req model.UpdateApprovalRequest) (model.UpdateApprovalResponse, error)
}

type ApprovalService struct {
	db     *sql.DB
	value  string
	logger log.Logger
}

// NewService makes a new Service.
func NewService(db *sql.DB, logger log.Logger) *ApprovalService {
	return &ApprovalService{
		db:     db,
		logger: logger,
	}
}

// Status only tell us that our service is ok!
func (ApprovalService) Status(ctx context.Context) (string, error) {
	return "ok", nil
}

// Get will return all approvals
func (ApprovalService) GetApprovals(ctx context.Context, req model.GetApprovalsRequest, db *sql.DB, logger log.Logger) (resp model.GetApprovalsResponse, err error) {
	return approval_logic.GetApprovals(ctx, req, db, logger)
}

// Add will add a new approval
func (ApprovalService) AddApproval(ctx context.Context, req model.AddApprovalRequest, db *sql.DB, logger log.Logger) (resp model.AddApprovalResponse, err error) {
	return approval_logic.AddApproval(ctx, req, db, logger)
}

// Update will update an existing approval
func (ApprovalService) UpdateApproval(ctx context.Context, req model.UpdateApprovalRequest, db *sql.DB, logger log.Logger) (resp model.UpdateApprovalResponse, err error) {
	return approval_logic.UpdateApproval(ctx, req, db, logger)
}
