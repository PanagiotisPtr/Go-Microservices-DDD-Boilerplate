package response

import (
	"context"
	"encoding/json"
	"net/http"
)

type GetAccountResponse struct {
	Uuid     string `json:"uuid"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func EncodeGetAccountResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
