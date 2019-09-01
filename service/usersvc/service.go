package usersvc

import (
	"context"
	"fmt"
	"stash-mono-repo/service/usersvc/model"
	"strconv"
)

// Service provides some "date capabilities" to your application
type Service interface {
	Status(ctx context.Context) (string, error)
	GetApproval(ctx context.Context, req model.GetApprovalRequest) (model.GetApprovalResponse, error)
}

type UserService struct{}

// NewService makes a new Service.
func NewService() *UserService {
	return &UserService{}
}

// Status only tell us that our service is ok!
func (UserService) Status(ctx context.Context) (string, error) {
	return "ok", nil
}

// Validate will check if the date today's date
func (UserService) GetApproval(ctx context.Context, req model.GetApprovalRequest) (resp model.GetApprovalResponse, err error) {
	resp = model.GetApprovalResponse{
		Status: "OK",
	}

	fmt.Println("Approval Received for " + req.ID + " Status: " + strconv.Itoa(req.Status))
	return
}
