package request

import (
	"context"
	"net/http"
)

type GetJwtRequest struct {
	RefreshToken string
}

func DecodeGetJwtRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req GetJwtRequest

	cookie, err := r.Cookie("RefreshToken")

	if err != nil {
		return nil, err
	}

	req.RefreshToken = cookie.Value

	return req, nil
}
