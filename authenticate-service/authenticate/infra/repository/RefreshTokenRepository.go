package repository

import (
	"authenticate-service/authenticate/domain/dto"
	"authenticate-service/authenticate/domain/irepository"
	"authenticate-service/authenticate/infra/entity"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/gofrs/uuid"

	"gorm.io/gorm"
)

type RefreshTokenRepository struct {
	db     *gorm.DB
	logger log.Logger
}

func NewRefreshTokenRepository(db *gorm.DB, logger log.Logger) irepository.IRefreshTokenRepository {
	return &RefreshTokenRepository{
		db:     db,
		logger: log.With(logger, "RefreshTokenRepository", "PostgreSQL"),
	}
}

func refreshTokenEntityToDto(etoken entity.RefreshToken) dto.RefreshTokenDto {
	return dto.RefreshTokenDto{
		Uuid:       etoken.UserUuid,
		UserUuid:   etoken.UserUuid,
		Expiration: etoken.Expiration,
	}
}

func (repo *RefreshTokenRepository) CreateToken(userUuid string) (dto.RefreshTokenDto, error) {
	var newToken entity.RefreshToken

	tokenUuid, err := uuid.NewV4()
	if err != nil {
		return refreshTokenEntityToDto(newToken), err
	}

	newToken = entity.RefreshToken{
		Uuid:       tokenUuid.String(),
		UserUuid:   userUuid,
		Expiration: time.Now().Unix(),
	}

	result := repo.db.Scopes(entity.RefreshTokenTable(newToken)).Create(&newToken)
	if result.Error != nil {
		return refreshTokenEntityToDto(newToken), result.Error
	}

	return refreshTokenEntityToDto(newToken), nil
}

func (repo *RefreshTokenRepository) GetToken(uuid string) (dto.RefreshTokenDto, error) {
	var token entity.RefreshToken
	token.Uuid = uuid
	result := repo.db.Scopes(entity.RefreshTokenTable(token)).First(&token)

	if result.Error != nil {
		return dto.RefreshTokenDto{}, result.Error
	}

	return refreshTokenEntityToDto(token), nil
}

func (repo *RefreshTokenRepository) DeleteToken(uuid string) error {
	var token entity.RefreshToken
	token.Uuid = uuid

	result := repo.db.Scopes(entity.RefreshTokenTable(token)).Delete(&token)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (repo *RefreshTokenRepository) RevokeUserTokens(userUuid string) error {
	var token entity.RefreshToken
	token.UserUuid = userUuid
	result := repo.db.Scopes(entity.RefreshTokenTable(token)).Delete(&token)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
