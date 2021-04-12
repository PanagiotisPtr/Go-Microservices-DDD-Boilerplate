package irepository

import (
	"authenticate-service/authenticate/domain/dto"
)

type IUserRepository interface {
	CreateUser(email string, password string) (dto.UserDto, error)
	GetUserByUuid(uuid string) (dto.UserDto, error)
	UpdateUser(user dto.UserDto) (dto.UserDto, error)
	DeleteUser(uuid string) error
	UserExists(email string) (bool, error)
	GetUserByEmail(email string) (dto.UserDto, error)
}
