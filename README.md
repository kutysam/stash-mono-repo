# Introduction
This repo is a mono repo and consists of 2 services [UserService and ApprovalService]. It is written in golang with go-kit microservices framework.
The purpose of this is to simulate a simple approval service based on some rules.
Current it is selfhosted on my Raspberry pi. Postgres, golang are all installed on it.
Use the swagger documentation if you want to test the APIs on the raspberry pi without going through local setup.
If it is down, please ping me.

# Local Setup
1. Ensure that you have golang installed
2. Clone the repo to $GOPATH/src e.g. /users/sathishramani/go/src/<REPO>
3. Make sure you have dep install. Otherwise, follow [instructions](https://golang.github.io/dep/docs/installation.html). After which, run `dep ensure`
4. Change the db connection string in main.go and create a DB called Approval and restore images/approvaldb.tar to have seed data. We are using postgresDB here.
5. `go run cmd/approvalsvc/main.go` [For approval service] && `go run cmd/usersvc/main.go` [For user service]
User service is a sample service created with service rule = 1. For objects that have been approved with service rule 1, they will be sent to userservice.

(I am using [go-kit](gokit.io) as the base of the microservice framework for this project)

# API Documentation
Refer to the swagger documentation [here](https://app.swaggerhub.com/apis/kutysam/StashApprovalAPI/1)
If you want to test, feel free to use the swagger UI. I've already appended the hosts to be kutysam.ddns.net and use click authorize first with `stashapprovalapikey` as the apikey.
To do a simple test on the service (Assuming all valid data),
1. PUT API (Create a new approval)
2. GET API (Get the request)
3. POST API (Change the status)

# How does the service work & Rules
## Working mechanism [Assuming all is valid]
1. See sequence diagram below.

## Rules
1. When an approval object is in ack / rejected / cancelled state, the user can ONLY edit the comment and nothing else.
2. When an approval object is in pending / error state, the user can once again either approve / reject / cancel the request.
3. When an approval object is in pending / error state, the user could also not edit the state and only edit the object information. When this happens, the object does not get sent to the designated service rule.
4. Error checking has been done on places where validation is required.

# Folder Structure
![FolderPath](/images/folderpath.png?raw=true "FolderPath")


# Database
Database is rather simple. We use this to store 2 main things
1. All approval objects
2. All service objects. A service rule object will contain the information on which server should the approval request be sent to given that the state has been changed to APPROVED / REJECTED / CANCELLED.
There is also a trigger which will auto update update_at column whenever the entry has been changed.
![Database](/images/db.png?raw=true "DB")


# Status Enum
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

![StateDiagram](/images/state.png?raw=true "State Diagram Image")

# Sequence Digram
There are only 4 endpoints. This sequence diagram is a high level overview on how the flow executes throughout the approval service.
Errors are checked accordingly but are not shown in the sequence diagram for now.
![SequenceDiagram](/images/sequence.png?raw=true "State Diagram Image")

# Additional Notes
1. Error checking has been done locally. Basically, I run through the flow using my own tests to ensure that the user cannot send invalid requests.
2. One important thing to take note in the logic is that, when a request confirmed to be approved / rejected / cancelled, the user can ONLY change the comment and nothing else. Do try it out and you will receive an error message if you input other changes like title etc.
3. When we send the request to the other server once it is in pending state, we will send the request JSON in the specified example format given in swagger. It is the duty of that service, to have a mapping of the original UUID approval object (when they created it) and logic depends on them. We only send a standard approval object JSON to the server when we receive an Approved / Rejected / Cancelled status.
4. Yes, the repo is a mono repo but, it is as such just for demonstration. User service is considered to be the '3rd' party service. There is no documentation for service rule, if you want to create a new service rule, please update the database directly for now. Right now, we only have 2 service rules, service rule 1 = http://kutysam.ddns.net:8001/approval and service rule 2 = invalidurl (Just to show error)

# Future Changes
## V1.0 [We cannot launch without these 3 important features]
1. Have another table called history. We will log down every change that happens to any approval object to here, be it, comment, title etc. This will be used as an audit log.
2. Have proper security measures (API_KEY for services and JWT for user accounts)
3. Have another column called previous state so that the user will know which state was he in, in case the state changes to error. Whether we should strictly follow the previous state, needs a team discussion.
4. Have a proper state machine diagram rather than using 'hardcoded' rules.
5. Proper error codes instead of all being 500.

## V2.0 onwards
1. Support approval dependencies such that, approval a will need 2 different services to successfully approve it in order for it to proceed to approved state.
2. If we need to scale to a large scale, we may move this out to a proper pub/sub mechanism to allivate load as the approval service calls other services.