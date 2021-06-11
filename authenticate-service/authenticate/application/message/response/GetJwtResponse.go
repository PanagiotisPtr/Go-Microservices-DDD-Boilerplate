package response

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

type GetJwtResponse struct {
	Success bool   `json:"success"`
	Token   string `json:"token"`
}

func EncodeGetJwtResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	getJwtResponse, ok := response.(GetJwtResponse)

	if ok == false {
		return errors.New("Response object could not be converted to AuthenticateUserResponse")
	}

	return json.NewEncoder(w).Encode(getJwtResponse)
}
