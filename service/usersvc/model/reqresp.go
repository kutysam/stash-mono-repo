package model

type GetApprovalRequest struct {
	ID     string
	Status int `json:"status,omitempty"`
}

type GetApprovalResponse struct {
	Status string `json:"status"`
}

type StatusRequest struct{}

type StatusResponse struct {
	Status string `json:"status"`
}
