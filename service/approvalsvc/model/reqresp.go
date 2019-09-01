package model

import "time"

//Healthcheck
type StatusRequest struct{}

type StatusResponse struct {
	Status string `json:"status"`
}

type GetApprovalsRequest struct {
	Limit    string            `json:"limit,omitempty"`
	Offset   string            `json:"offset,omitempty"`
	Restrict map[string]string `json:"restrict,omitempty"`
	Sort     map[string]string `json:"sort,omitempty"`
}

type GetApprovalsResponse struct {
	ApprovalItems []ApprovalItem `json:"approval_items"`
	Err           string         `json:"err,omitempty"`
}

// Add a new request
type AddApprovalRequest struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	ServiceRule int       `json:"service_rule"`
	Deadline    time.Time `json:"deadline,omitempty"`
	Comment     string    `json:"comment,omitempty"`
}

type AddApprovalResponse struct {
	ApprovalItem ApprovalItem `json:"approval_item,omitempty"`
	Err          string       `json:"err,omitempty"`
}

// Update the approval request
type UpdateApprovalRequest struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	ServiceRule int       `json:"service_rule"`
	Deadline    time.Time `json:"deadline,omitempty"`
	Comment     string    `json:"comment,omitempty"`
	Status      int       `json:"status,omitempty"`
}

// Change priority / status etc.
type UpdateApprovalResponse struct {
	ApprovalItem ApprovalItem `json:"approval_item"`
	Err          string       `json:"err"`
}
