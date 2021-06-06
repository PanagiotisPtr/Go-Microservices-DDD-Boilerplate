package repository

import (
	"authenticate-service/authenticate/domain/dto"
	"authenticate-service/authenticate/domain/irepository"
	"authenticate-service/authenticate/infra/entity"

	"github.com/go-kit/kit/log"
	"github.com/gofrs/uuid"

	"errors"

	"gorm.io/gorm"
)

type UserRepository struct {
	db     *gorm.DB
	logger log.Logger
}

func userEntityToDto(euser entity.User) dto.UserDto {
	return dto.UserDto{
		Uuid:     euser.Uuid,
		Email:    euser.Email,
		Password: euser.Password,
		Verified: euser.Verified,
	}
}

func NewUserRepository(db *gorm.DB, logger log.Logger) irepository.IUserRepository {
	return &UserRepository{
		db:     db,
		logger: log.With(logger, "UserRepository", "PostgreSQL"),
	}
}

func (repo *UserRepository) CreateUser(email string, password string) (dto.UserDto, error) {
	var user entity.User

	uuid, err := uuid.NewV4()
	if err != nil {
		return userEntityToDto(user), err
	}

	user = entity.User{
		Uuid:     uuid.String(),
		Email:    email,
		Password: password,
		Verified: false,
	}

	result := repo.db.Scopes(entity.UserTable(user)).Create(&user)
	if result.Error != nil {
		return userEntityToDto(user), result.Error
	}

	return userEntityToDto(user), nil
}

func (repo *UserRepository) GetUserByUuid(uuid string) (dto.UserDto, error) {
	var user entity.User
	user.Uuid = uuid
	result := repo.db.Scopes(entity.UserTable(user)).First(&user)

	if result.Error != nil {
		return userEntityToDto(user), result.Error
	}

	return userEntityToDto(user), nil
}

func (repo *UserRepository) UpdateUser(userDto dto.UserDto) (dto.UserDto, error) {
	var user entity.User
	user.Uuid = userDto.Uuid

	result := repo.db.Scopes(entity.UserTable(user)).First(&user)
	if result.Error != nil {
		return userEntityToDto(user), result.Error
	}

	user.Email = userDto.Email
	user.Password = userDto.Password
	user.Verified = userDto.Verified

	result = repo.db.Scopes(entity.UserTable(user)).Save(&user)
	if result.Error != nil {
		return userEntityToDto(user), result.Error
	}

	return userEntityToDto(user), nil
}

func (repo *UserRepository) DeleteUser(uuid string) error {
	var user entity.User
	user.Uuid = uuid

	result := repo.db.Scopes(entity.UserTable(user)).Delete(&user)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (repo *UserRepository) GetUserByEmail(email string) (dto.UserDto, error) {
	var user entity.User
	user.Email = email

	result := repo.db.Scopes(entity.UserTable(user)).First(&user)

	return userEntityToDto(user), result.Error
}

func (repo *UserRepository) UserExists(email string) (bool, error) {
	var user entity.User
	user.Email = email

	result := repo.db.Scopes(entity.UserTable(user)).First(&user)

	if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return false, result.Error
	}

	return result.RowsAffected > 0, nil
}
