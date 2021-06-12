package request

import (
	"context"
	"net/http"
)

type LogoutUserRequest struct {
	RefreshToken string
}

func DecodeLogoutUserRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req LogoutUserRequest
	cookie, err := r.Cookie("RefreshToken")

	if err != nil {
		return nil, err
	}

	req.RefreshToken = cookie.Value

	return req, nil
}
