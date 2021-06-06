package request

import (
	"context"
	"encoding/json"
	"net/http"
)

type AuthenticateUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func DecodeAuthenticateUserRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req AuthenticateUserRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}

	return req, nil
}
