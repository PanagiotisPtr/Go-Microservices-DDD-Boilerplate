package response

import (
	"context"
	"encoding/json"
	"net/http"
)

type RegisterUserResponse struct {
	Success bool `json:"success"`
}

func EncodeRegisterUserResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
