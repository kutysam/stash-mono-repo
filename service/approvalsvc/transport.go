package approvalsvc

import (
	"context"
	"encoding/json"
	"net/http"
	"stash-mono-repo/service/approvalsvc/model"
	"strconv"
)

// In the first part of the file we are mapping requests and responses to their JSON payload.

// In the second part we will write "decoders" for our incoming requests
func decodeGetApprovalRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req model.GetApprovalRequest
	keys, ok := r.URL.Query()["id"]
	if ok {
		req.ID = keys[0]
	}

	return req, nil
}

func decodeGetApprovalsRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req model.GetApprovalsRequest
	keys, ok := r.URL.Query()["limit"]
	if ok {
		defaultInt, err := strconv.Atoi(keys[0])
		if err != nil {
			req.Limit = -1
		}
		req.Limit = defaultInt
	} else {
		req.Limit = -1
	}

	keys, ok = r.URL.Query()["offset"]
	if ok {
		defaultInt, err := strconv.Atoi(keys[0])
		if err != nil {
			req.Offset = -1
		}
		req.Offset = defaultInt
	} else {
		req.Offset = -1
	}

	keys, ok = r.URL.Query()["default"]
	if ok {
		defaultInt, err := strconv.Atoi(keys[0])
		if err != nil {
			req.Default = 0
		}
		req.Default = defaultInt
	} else {
		req.Default = 0
	}

	return req, nil
}

func decodeAddApprovalRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req model.AddApprovalRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func decodeUpdateApprovalRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req model.UpdateApprovalRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func decodeStatusRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req model.StatusRequest
	return req, nil
}

// Last but not least, we have the encoder for the response output
func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
