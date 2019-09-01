package approvalsvc

import (
	"context"
	"stash-mono-repo/service/approvalsvc/model"

	"github.com/go-kit/kit/endpoint"
)

// Endpoints are exposed
type Endpoints struct {
	StatusEndpoint         endpoint.Endpoint
	GetApprovalEndpoint    endpoint.Endpoint
	GetApprovalsEndpoint   endpoint.Endpoint
	AddApprovalEndpoint    endpoint.Endpoint
	UpdateApprovalEndpoint endpoint.Endpoint
}

// MakeStatusEndpoint returns the response from our service "status"
func MakeStatusEndpoint(srv ApprovalService) endpoint.Endpoint {
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
func MakeGetApprovalEndpoint(srv ApprovalService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(model.GetApprovalRequest)
		d, err := srv.GetApproval(ctx, req, srv.db, srv.logger)
		return d, err
	}
}

// GetApprovals endpoint mapping
func (e Endpoints) GetApproval(ctx context.Context) (model.GetApprovalResponse, error) {
	req := model.GetApprovalsRequest{}
	resp, err := e.GetApprovalEndpoint(ctx, req)
	if err != nil {
		return model.GetApprovalResponse{}, err
	}
	getApprovalsResp := resp.(model.GetApprovalResponse)
	return getApprovalsResp, nil
}

// MakeGetApprovalsEndpoint returns the response from our service "get"
func MakeGetApprovalsEndpoint(srv ApprovalService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(model.GetApprovalsRequest)
		d, err := srv.GetApprovals(ctx, req, srv.db, srv.logger)
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
func MakeAddApprovalEndpoint(srv ApprovalService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(model.AddApprovalRequest)
		d, err := srv.AddApproval(ctx, req, srv.db, srv.logger)
		return d, err
	}
}

// AddApproval endpoint mapping
func (e Endpoints) AddApproval(ctx context.Context) (model.AddApprovalResponse, error) {
	req := model.AddApprovalRequest{}
	resp, err := e.AddApprovalEndpoint(ctx, req)
	if err != nil {
		return model.AddApprovalResponse{}, err
	}

	addApprovalsResp := resp.(model.AddApprovalResponse)
	return addApprovalsResp, nil
}

// MakeUpdateRequestEndpoint returns the response from our service "put"
func MakeUpdateApprovalEndpoint(srv ApprovalService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(model.UpdateApprovalRequest)
		d, err := srv.UpdateApproval(ctx, req, srv.db, srv.logger)
		return d, err
	}
}

// AUpdateApprovaldApproval endpoint mapping
func (e Endpoints) UpdateApproval(ctx context.Context) (model.UpdateApprovalResponse, error) {
	req := model.UpdateApprovalRequest{}
	resp, err := e.UpdateApprovalEndpoint(ctx, req)
	if err != nil {
		return model.UpdateApprovalResponse{}, err
	}
	updateApprovalsResp := resp.(model.UpdateApprovalResponse)
	return updateApprovalsResp, nil
}
