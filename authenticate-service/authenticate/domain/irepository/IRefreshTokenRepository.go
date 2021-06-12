package irepository

import (
	"authenticate-service/authenticate/domain/dto"
)

type IRefreshTokenRepository interface {
	CreateToken(userUuid string) (dto.RefreshTokenDto, error)
	GetToken(uuid string) (dto.RefreshTokenDto, error)
	DeleteToken(uuid string) error
	RevokeUserTokens(refreshTokenUuid string) error
}
