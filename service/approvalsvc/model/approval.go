package model

import (
	"time"
)

type ApprovalItem struct {
	ID          string    `json:"id"`
	ServiceRule int       `json:"service_rule,omitempty"`
	Priority    int       `json:"priority,omitempty"`
	Deadline    time.Time `json:"deadline,omitempty"`
	Comment     string    `json:"comment,omitempty"`
	Status      string    `json:"status,omitempty"`
}
