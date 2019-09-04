# Introduction
This repo is a mono repo and consists of 2 services [UserService and ApprovalService]. It is written in golang with go-kit microservices framework.

The purpose of this is to simulate a simple approval service based on some rules.

# Installation
1. Ensure that you have golang installed
2. Clone the repo to $GOPATH/src e.g. /users/sathishramani/go/src/<REPO>
3. `go run cmd/approvalsvc/main.go` [For approval service] && `go run cmd/usersvc/main.go` [For user service]
4. [Optional], change the db connection string in main.go and run the SQL script `approvaldb.sql` to have seed data. We are using postgresDB here.

# API Documentation
Refer to the swagger documentation [here](https://app.swaggerhub.com/apis/kutysam/StashApprovalAPI/1)

# Enums and description

## Status Enum
Refer to service\approvalsvc\model\approval.go

| Enum  | Status | Description |
| ------------- | ------------- | ------------- |
| 0 | STATUS_UNKNOWN | This enum should never ever be reached. If somehow this enum is reached, investigation must be done. |
| 1 | STATUS_ERROR | Error happened when we are calling the service once the approver has requested this approval item to be either 3,4,5. The approver can redo their operation and choose another operation instead of the one they selected before. [See sequence diagram for a better example] |
| 2 | STATUS_PENDING | Every new request will be created with PENDING status. |
| 3 | STATUS_APPROVED | This approval item has been successfully approved. This will only be reached after state 6. |
| 4 | STATUS_REJECTED | Similar to 3 |
| 5 | STATUS_CANCELLED | Similar to 3 |
| 6 | STATUS_ACKNOWLEDGED_APPROVED | When the approver submits their request, before making a call to the designated server, the object will be marked as this state if the approver marks it as approved  |
| 7 | STATUS_ACKNOWLEDGED_REJECTED | Similar to 6 |
| 8 | STATUS_ACKNOWLEDGED_CANCELLED | Similar to 6 |