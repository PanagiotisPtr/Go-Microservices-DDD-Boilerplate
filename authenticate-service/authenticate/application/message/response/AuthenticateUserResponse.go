package response

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

type AuthenticateUserResponse struct {
	Success      bool   `json:"success"`
	Token        string `json:"token"`
	RefreshToken string `json:"-"` // important: keep this as '-'
}

func EncodeAuthenticateUserResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	authenticateUserResponse, ok := response.(AuthenticateUserResponse)

	if ok == false {
		return errors.New("Response object could not be converted to AuthenticateUserResponse")
	}

	if authenticateUserResponse.RefreshToken == "" {
		return errors.New("Response is missing refresh token")
	}

	cookie := http.Cookie{
		Name:     "RefreshToken",
		Value:    authenticateUserResponse.RefreshToken,
		Path:     "/",
		Expires:  time.Now().Add(time.Hour*24 - time.Second),
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
	}

	http.SetCookie(w, &cookie)

	return json.NewEncoder(w).Encode(response)
}
