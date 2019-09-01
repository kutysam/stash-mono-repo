package usersvc

import (
	"context"
	"stash-mono-repo/service/usersvc/model"

	"github.com/go-kit/kit/endpoint"
)

// Endpoints are exposed
type Endpoints struct {
	StatusEndpoint   endpoint.Endpoint
	ApprovalEndpoint endpoint.Endpoint
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

// MakeApprovalEndpoint returns the response from our service "approval"
func MakeApprovalEndpoint(srv Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(model.GetApprovalRequest)
		d, err := srv.GetApproval(ctx, req)
		return d, err
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

// Approval endpoint mapping
func (e Endpoints) GetApproval(ctx context.Context, date string) (model.GetApprovalResponse, error) {
	req := model.GetApprovalRequest{}
	resp, err := e.ApprovalEndpoint(ctx, req)
	if err != nil {
		return model.GetApprovalResponse{}, err
	}
	approvalResp := resp.(model.GetApprovalResponse)
	return approvalResp, nil
}
