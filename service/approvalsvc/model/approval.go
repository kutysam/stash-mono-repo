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

const (
	STATUS_UNKNOWN   = 0
	STATUS_PENDING   = 1
	STATUS_APPROVED  = 2
	STATUS_REJECTED  = 3
	STATUS_CANCELLED = 4
)
