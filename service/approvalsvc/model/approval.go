package model

import (
	"time"
)

type ApprovalItem struct {
	ID          string     `json:"id,omitempty"`
	ServiceRule int        `json:"service_rule,omitempty"`
	Priority    int        `json:"priority,omitempty"`
	Deadline    *time.Time `json:"deadline,omitempty"`
	Status      int        `json:"status,omitempty"`
	Comment     string     `json:"comment,omitempty"`
}

type Priority int

const (
	PRIORITY_DEFAULT = 0
	PRIORITY_HIGHEST = 1
	PRIORITY_HIGH    = 2
	PRIORITY_MEDIUM  = 3
	PRIORITY_LOW     = 4
	PRIORITY_LOWEST  = 5

	STATUS_UNKNOWN   = 0
	STATUS_PENDING   = 1
	STATUS_APPROVED  = 2
	STATUS_REJECTED  = 3
	STATUS_CANCELLED = 4
)

var ServiceRule = map[int][]string{
	1: []string{"http://localhost:8000/usersvc/getapproval", "testapikey"},
	2: []string{"http", ""},
}
