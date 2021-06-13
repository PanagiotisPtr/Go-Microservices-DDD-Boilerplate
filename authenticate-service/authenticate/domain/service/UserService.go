package service

import (
	"authenticate-service/authenticate/domain/generator"
	"authenticate-service/authenticate/domain/irepository"
	"errors"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepo       irepository.IUserRepository
	refreshRepo    irepository.IRefreshTokenRepository
	logger         log.Logger
	tokenGenerator generator.JwtGenerator
}

func NewUserService(userRepo irepository.IUserRepository, refreshRepo irepository.IRefreshTokenRepository, logger log.Logger, tokenGenerator generator.JwtGenerator) UserService {
	return UserService{
		userRepo:       userRepo,
		refreshRepo:    refreshRepo,
		logger:         logger,
		tokenGenerator: tokenGenerator,
	}
}

func (s *UserService) RegisterUser(email string, password string) error {
	if len(password) < 5 {
		return errors.New("Password too short")
	}

	if email == "" {
		return errors.New("Missing email")
	}

	userExists, err := s.userRepo.UserExists(email)
	if err != nil {
		return err
	}

	if userExists {
		return errors.New("User with this email already exists")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user, err := s.userRepo.CreateUser(email, string(hash))

	level.Info(s.logger).Log("Create user with", "uuid", user.Uuid)

	return nil
}

func (s *UserService) AuthenticateUser(email string, password string) (string, error) {
	user, err := s.userRepo.GetUserByEmail(email)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", err
	}

	tokenDto, err := s.refreshRepo.CreateToken(user.Uuid)
	if err != nil {
		return "", err
	}

	refreshToken, err := s.tokenGenerator.CreateToken(time.Hour*24, user.Uuid, tokenDto.Uuid)
	if err != nil {
		return "", err
	}

	return refreshToken, nil
}

func (s *UserService) GetJWT(refreshToken string) (string, error) {
	claims, err := s.tokenGenerator.ValidateToken(refreshToken)
	if err != nil {
		return "", err
	}

	userUuid, ok := claims["sub"].(string)
	if !ok {
		return "", errors.New("Missing user uuid in refresh token claims")
	}

	tokenUuid, ok := claims["dat"].(string)
	if !ok {
		return "", errors.New("Missing user uuid in refresh token claims")
	}

	tokenDto, err := s.refreshRepo.GetToken(tokenUuid)
	if err != nil {
		return "", err
	}

	if tokenDto.Expiration < time.Now().UTC().Unix() {
		if err := s.refreshRepo.DeleteToken(tokenUuid); err != nil {
			return "", err
		}

		return "", errors.New("Refresh token has expired")
	}

	if tokenDto.UserUuid != userUuid {
		return "", errors.New("Refresh token doesn't match user")
	}

	user, err := s.userRepo.GetUserByUuid(userUuid)
	if err != nil {
		return "", err
	}

	newToken, err := s.tokenGenerator.CreateToken(time.Minute*15, user.Uuid, user.Email)

	return newToken, nil
}

func (s *UserService) GetUserUuidFromToken(jwtToken string) (string, error) {
	claims, err := s.tokenGenerator.ValidateToken(jwtToken)
	if err != nil {
		return "", err
	}

	userUuid, ok := claims["sub"].(string)

	if ok == false {
		return "", errors.New("JWT is missing the user uuid")
	}

	return userUuid, nil
}

func (s *UserService) VerifyUser(uuid string, code string) error {
	// check if verification code is correct

	return nil
}

func (s *UserService) UpdateUser(
	userUuid string,
	oldPassword string,
	newPassword string,
) error {
	user, err := s.userRepo.GetUserByUuid(userUuid)

	if err != nil {
		return err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(oldPassword))
	if err != nil {
		return err
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = string(hash)

	_, err = s.userRepo.UpdateUser(user)

	return err
}

func (s *UserService) RevokeRefreshTokensForUser(token string) error {
	claims, err := s.tokenGenerator.ValidateToken(token)
	if err != nil {
		return err
	}

	refreshTokenUuid, ok := claims["dat"].(string)
	if !ok {
		return errors.New("Missing user uuid in refresh token claims")
	}

	return s.refreshRepo.RevokeUserTokens(refreshTokenUuid)
}
