package usersvc

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"stash-mono-repo/service/usersvc/model"
)

const (
	apikey        = "usersvcapikey" //TODO: Make it secure
	invalidApiKey = "Invalid Api Key Provided"
)

// We will write "decoders" for our incoming requests

func decodeApprovalRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	if !checkAuthentication(r) {
		return nil, errors.New(invalidApiKey)
	}
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

func checkAuthentication(r *http.Request) bool {
	if val, ok := r.Header["Authorization"]; ok {
		if val[0] != apikey {
			return false
		}
	} else {
		return false
	}
	return true
}
