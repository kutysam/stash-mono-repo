package usersvc

import (
	"context"
	"encoding/json"
	"net/http"
	"stash-mono-repo/service/usersvc/model"
)

// In the first part of the file we are mapping requests and responses to their JSON payload.

// In the second part we will write "decoders" for our incoming requests

func decodeApprovalRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req model.GetApprovalRequest
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
