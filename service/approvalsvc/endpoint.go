package approvalsvc

import (
	"context"
	"stash-mono-repo/service/approvalsvc/model"

	"github.com/go-kit/kit/endpoint"
)

// Endpoints are exposed
type Endpoints struct {
	StatusEndpoint         endpoint.Endpoint
	GetApprovalsEndpoint   endpoint.Endpoint
	AddApprovalEndpoint    endpoint.Endpoint
	UpdateApprovalEndpoint endpoint.Endpoint
}

// MakeStatusEndpoint returns the response from our service "status"
func MakeStatusEndpoint(srv Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		_ = request.(model.StatusRequest) // we really just need the request, we don't use any value from it
		s, err := srv.Status(ctx)
		if err != nil {
			return model.StatusResponse{s}, err
		}
		return model.StatusResponse{s}, nil
	}
}

// Status endpoint mapping
func (e Endpoints) Status(ctx context.Context) (string, error) {
	req := model.StatusRequest{}
	resp, err := e.StatusEndpoint(ctx, req)
	if err != nil {
		return "", err
	}
	statusResp := resp.(model.StatusResponse)
	return statusResp.Status, nil
}

// MakeGetApprovalsEndpoint returns the response from our service "get"
func MakeGetApprovalsEndpoint(srv Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(model.GetApprovalsRequest)
		d, err := srv.GetApprovals(ctx, req)
		return d, err
	}
}

// GetApprovals endpoint mapping
func (e Endpoints) GetApprovals(ctx context.Context) (model.GetApprovalsResponse, error) {
	req := model.GetApprovalsRequest{}
	resp, err := e.GetApprovalsEndpoint(ctx, req)
	if err != nil {
		return model.GetApprovalsResponse{}, err
	}
	getApprovalsResp := resp.(model.GetApprovalsResponse)
	return getApprovalsResp, nil
}

// MakeAddRequestEndpoint returns the response from our service "post"
func MakeAddApprovalEndpoint(srv Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(model.AddApprovalRequest)
		d, err := srv.AddApproval(ctx, req)
		return d, err
	}
}

// AddApproval endpoint mapping
func (e Endpoints) AddApproval(ctx context.Context, a model.ApprovalItem) (model.AddApprovalResponse, error) {
	req := model.AddApprovalRequest{
		ID:          a.ID,
		ServiceRule: a.ServiceRule,
		Priority:    a.Priority,
		Deadline:    a.Deadline,
		Comment:     a.Comment,
	}

	resp, err := e.AddApprovalEndpoint(ctx, req)
	if err != nil {
		return model.AddApprovalResponse{}, err
	}

	addApprovalsResp := resp.(model.AddApprovalResponse)
	return addApprovalsResp, nil
}

// MakeUpdateRequestEndpoint returns the response from our service "put"
func MakeUpdateApprovalEndpoint(srv Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(model.UpdateApprovalRequest)
		d, err := srv.UpdateApproval(ctx, req)
		return d, err
	}
}

// AUpdateApprovaldApproval endpoint mapping
func (e Endpoints) UpdateApproval(ctx context.Context) (model.UpdateApprovalResponse, error) {
	req := model.AddApprovalRequest{}
	resp, err := e.UpdateApprovalEndpoint(ctx, req)
	if err != nil {
		return model.UpdateApprovalResponse{}, err
	}
	updateApprovalsResp := resp.(model.UpdateApprovalResponse)
	return updateApprovalsResp, nil
}
