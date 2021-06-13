package request

import (
	"context"
	"encoding/json"
	"net/http"
)

type UpdateUserRequest struct {
	RefreshToken string
	OldPassword  string `json:"oldPassword"`
	NewPassword  string `json:"newPassword"`
}

func DecodeUpdateUserRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req UpdateUserRequest
	cookie, err := r.Cookie("RefreshToken")

	if err != nil {
		return nil, err
	}

	err = json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		return nil, err
	}

	req.RefreshToken = cookie.Value

	return req, nil
}
