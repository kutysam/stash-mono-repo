package model

type GetApprovalRequest struct {
	ID     string `json:"id"`
	Status int    `json:"status"`
}

type GetApprovalResponse struct {
	ID        string `json: "id"`
	Status    int    `json:"status"`
	ErrorCode int    `json:"error_code"`
}

type StatusRequest struct{}

type StatusResponse struct {
	Status string `json:"status"`
}
