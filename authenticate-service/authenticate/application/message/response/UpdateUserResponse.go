package response

import (
	"context"
	"encoding/json"
	"net/http"
)

type UpdateUserResponse struct {
	Success bool `json:"success"`
}

func EncodeUpdateUserResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
