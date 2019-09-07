package model

import (
	"time"
)

type ApprovalItem struct {
	ID          string     `json:"id,omitempty"`
	Title       string     `json:"title,omitempty"`
	Description string     `json:"description,omitempty"`
	ServiceRule int        `json:"service_rule,omitempty"`
	Deadline    *time.Time `json:"deadline,omitempty"`
	Status      int        `json:"status,omitempty"`
	Comment     string     `json:"comment,omitempty"`
	CreatedAt   *time.Time `json:"created_at,omitempty"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
}

type ServiceRule struct {
	Name   string
	Apikey string
	URL    string
}

const (
	STATUS_UNKNOWN                = 0
	STATUS_ERROR                  = 1
	STATUS_PENDING                = 2
	STATUS_APPROVED               = 3
	STATUS_REJECTED               = 4
	STATUS_CANCELLED              = 5
	STATUS_ACKNOWLEDGED_APPROVED  = 6 //Server Only Option
	STATUS_ACKNOWLEDGED_REJECTED  = 7 //Server Only Option
	STATUS_ACKNOWLEDGED_CANCELLED = 8 //Server Only Option
	APPROVAL_TABLE                = "approval"
	SERVICE_RULE_TABLE            = "servicerule"
)

func CheckValidStatus(status int) bool {
	if status >= STATUS_APPROVED && status <= STATUS_CANCELLED {
		return true
	} else {
		return false
	}
}
