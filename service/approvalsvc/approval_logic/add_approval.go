package approval_logic

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"stash-mono-repo/service/approvalsvc/model"
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

const (
	sqlInsertStatement = `
	INSERT INTO approval (id, title, description, service_rule,comment, status, deadline, created_at, updated_at)
	VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9)`
)

func AddApproval(ctx context.Context, req model.AddApprovalRequest, db *sql.DB, logger log.Logger) (resp model.AddApprovalResponse, err error) {
	id := uuid.New().String()
	title := req.Title
	description := req.Description
	serviceRule := req.ServiceRule
	deadline := &req.Deadline
	comment := req.Comment
	createdAt := time.Now()
	status := model.STATUS_PENDING

	invalidFields := ""
	if title == "" {
		invalidFields += " Title is empty."
	}

	if description == "" {
		invalidFields += " Description is empty."
	}

	if invalidFields != "" {
		err = errors.New(model.ERROR_INVALID_FIELDS + invalidFields)
		return
	}

	if time.Now().After(*deadline) {
		err = errors.New(model.ERROR_CURRENT_TIME_IS_AFTER_DEADLINE)
		return
	}

	approvalItem := model.ApprovalItem{
		ID:          id,
		Title:       title,
		Description: description,
		ServiceRule: serviceRule,
		Comment:     comment,
		Status:      model.STATUS_PENDING,
		Deadline:    deadline,
		CreatedAt:   &createdAt,
	}

	_, err = db.Exec(sqlInsertStatement, id, title, description, serviceRule, comment, status, deadline, createdAt, createdAt)
	if err != nil {
		pqErr := err.(*pq.Error)
		if pqErr.Code == model.PQ_ERROR_FOREIGN_KEY_VIOLATION && pqErr.Constraint == model.PQ_SERVICE_RULE_CONSTRAINT {
			err = errors.New(model.ERROR_SERVICE_RULE_INVALID_CONSTRAINT)
			return
		}
		err = errors.New(model.ERROR_DATABASE_ERROR)
		return
	}

	resp = model.AddApprovalResponse{
		ApprovalItem: approvalItem,
	}

	return
}
