package generator

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JwtGenerator struct {
	privateKey []byte
	publicKey  []byte
}

func NewJwtGenerator(privateKey []byte, publicKey []byte) JwtGenerator {
	return JwtGenerator{
		privateKey: privateKey,
		publicKey:  publicKey,
	}
}

func (generator *JwtGenerator) CreateToken(duration time.Duration, subject string, data interface{}) (string, error) {
	key, err := jwt.ParseRSAPrivateKeyFromPEM(generator.privateKey)
	if err != nil {
		return "", err
	}

	now := time.Now().UTC()

	claims := make(jwt.MapClaims)
	claims["sub"] = subject
	claims["dat"] = data
	claims["exp"] = now.Add(duration).Unix()

	token, err := jwt.NewWithClaims(jwt.SigningMethodRS256, claims).SignedString(key)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (generator *JwtGenerator) ValidateToken(token string) (jwt.MapClaims, error) {
	var claims jwt.MapClaims

	key, err := jwt.ParseRSAPublicKeyFromPEM(generator.publicKey)
	if err != nil {
		return claims, err
	}

	result, err := jwt.Parse(token, func(jwtToken *jwt.Token) (interface{}, error) {
		if _, ok := jwtToken.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, err
		}

		return key, nil
	})
	if err != nil {
		return claims, err
	}

	claims, ok := result.Claims.(jwt.MapClaims)
	if !ok || !result.Valid {
		return claims, errors.New("Invalid token")
	}

	exp, ok := claims["exp"].(float64)
	if !ok {
		return claims, errors.New("Incorrect token structure")
	}

	if int64(exp) < time.Now().UTC().Unix() {
		return claims, errors.New("Token has expired")
	}

	return claims, nil
}
