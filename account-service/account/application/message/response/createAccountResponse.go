package response

import (
	"context"
	"encoding/json"
	"net/http"
)

type CreateAccountResponse struct {
	Uuid string `json:"uuid"`
}

func EncodeCreateAccountResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
