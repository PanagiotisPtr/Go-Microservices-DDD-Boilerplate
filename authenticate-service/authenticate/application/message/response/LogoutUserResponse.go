package response

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

type LogoutUserResponse struct {
	Success bool `json:"success"`
}

func EncodeLogoutUserResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	logoutUserResponse, ok := response.(LogoutUserResponse)

	if ok == false {
		return errors.New("Response object could not be converted to LogoutUserResponse")
	}

	return json.NewEncoder(w).Encode(logoutUserResponse)
}
