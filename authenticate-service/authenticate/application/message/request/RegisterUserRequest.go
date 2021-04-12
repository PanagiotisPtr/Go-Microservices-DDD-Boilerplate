package request

import (
	"context"
	"encoding/json"
	"net/http"
)

type RegisterUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func DecodeRegisterUserRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req RegisterUserRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}

	return req, nil
}
