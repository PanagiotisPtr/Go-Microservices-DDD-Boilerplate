package response

import (
	"context"
	"encoding/json"
	"net/http"
)

type AuthenticateUserResponse struct {
	Success bool   `json:"success"`
	Token   string `json:"token"`
}

func EncodeAuthenticateUserResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
