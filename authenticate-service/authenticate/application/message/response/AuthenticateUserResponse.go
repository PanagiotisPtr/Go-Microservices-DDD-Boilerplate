package response

import (
	"context"
	"errors"
	"net/http"
	"time"
)

type AuthenticateUserResponse struct {
	Success bool   `json:"success"`
	Token   string `json:"token"`
}

func EncodeAuthenticateUserResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	authenticateUserResponse, ok := response.(AuthenticateUserResponse)

	if ok == false {
		return errors.New("Response object could not be converted to AuthenticateUserResponse")
	}

	if authenticateUserResponse.Token == "" {
		return errors.New("Response is missing refresh token")
	}

	cookie := http.Cookie{
		Name:     "RefreshToken",
		Value:    authenticateUserResponse.Token,
		Path:     "/",
		Expires:  time.Now().Add(time.Hour*24 - time.Second),
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
	}

	http.SetCookie(w, &cookie)

	return nil
}
