package request

import (
	"context"
	"encoding/json"
	"net/http"
)

type CreateAccountRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func DecodeCreateAccountRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req CreateAccountRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}

	return req, nil
}
