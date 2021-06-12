package application

import (
	applicationEndpoints "authenticate-service/authenticate/application/endpoint"
	"authenticate-service/authenticate/application/server"
	"authenticate-service/authenticate/domain/generator"
	"authenticate-service/authenticate/domain/service"
	"authenticate-service/authenticate/infra/entity"
	"authenticate-service/authenticate/infra/repository"
	"authenticate-service/config"
	"context"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"

	"github.com/go-kit/kit/log/level"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Bootstrap(ctx context.Context, logger log.Logger, configuration config.Configuration) http.Handler {
	// load private and public keys
	privateKey, err := ioutil.ReadFile(configuration.Service.PrivateKeyLocation)
	if err != nil {
		level.Error(logger).Log("exit", err)
		os.Exit(-1)
	}

	publicKey, err := ioutil.ReadFile(configuration.Service.PublicKeyLocation)
	if err != nil {
		level.Error(logger).Log("exit", err)
		os.Exit(-1)
	}

	// Connect to the database
	dsn := "host=" + configuration.Database.Host + " user=" + configuration.Database.Username +
		" password=" + configuration.Database.Password + " dbname=" + configuration.Database.Name +
		" port=" + configuration.Database.Port + " sslmode=disable TimeZone=Europe/London"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		level.Error(logger).Log("exit", err)
		os.Exit(-1)
	}

	// Migrate all entities (only one in this case)
	db.Scopes(entity.UserTable(entity.User{})).AutoMigrate(&entity.User{})
	db.Scopes(entity.RefreshTokenTable(entity.RefreshToken{})).AutoMigrate(&entity.RefreshToken{})

	jwtGenerator := generator.NewJwtGenerator(privateKey, publicKey)
	userRepository := repository.NewUserRepository(db, logger)
	refreshTokenRepository := repository.NewRefreshTokenRepository(db, logger)
	userService := service.NewUserService(
		userRepository,
		refreshTokenRepository,
		logger,
		jwtGenerator,
	)

	endpoints := make(map[string]endpoint.Endpoint)

	endpoints["RegisterUserEndpoint"] = applicationEndpoints.RegisterUserEndpoint(userService)
	endpoints["AuthenticateEndpoint"] = applicationEndpoints.AuthenticateUserEndpoint(userService)
	endpoints["GetJwtEndpoint"] = applicationEndpoints.GetJwtEndpoint(userService)
	endpoints["LogoutUserEndpoint"] = applicationEndpoints.LogoutUserEndpoint(userService)

	return server.NewHTTPServer(ctx, endpoints)
}
