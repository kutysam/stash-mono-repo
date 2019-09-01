package approvalsvc

import (
	"context"
	"encoding/json"
	"net/http"
	"stash-mono-repo/service/approvalsvc/model"
)

// In the first part of the file we are mapping requests and responses to their JSON payload.

// In the second part we will write "decoders" for our incoming requests
func decodeGetApprovalsRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req model.GetApprovalsRequest
	keys, ok := r.URL.Query()["limit"]
	if ok {
		req.Limit = keys[0]
	}

	keys, ok = r.URL.Query()["offset"]
	if ok {
		req.Offset = keys[0]
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
