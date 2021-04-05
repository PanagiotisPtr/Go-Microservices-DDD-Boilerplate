package application

import (
	applicationEndpoints "account-service/account/application/endpoint"
	"account-service/account/application/server"
	"account-service/account/domain/service"
	"account-service/account/infra/entity"
	"account-service/account/infra/repository"
	"account-service/config"
	"context"
	"net/http"
	"os"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"

	"github.com/go-kit/kit/log/level"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Bootstrap(ctx context.Context, logger log.Logger, configuration config.Configuration) http.Handler {
	dsn := "host=" + configuration.Database.Host + " user=" + configuration.Database.Username +
		" password=" + configuration.Database.Password + " dbname=" + configuration.Database.Name +
		" port=" + configuration.Database.Port + " sslmode=disable TimeZone=Europe/London"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		level.Error(logger).Log("exit", err)
		os.Exit(-1)
	}

	// Migrate all entities (only one in this case)
	db.Scopes(entity.AccountTable(entity.Account{})).AutoMigrate(&entity.Account{})

	repository := repository.NewAccountRepository(db, logger)
	accountService := service.NewAccountService(repository, logger)

	endpoints := make(map[string]endpoint.Endpoint)

	endpoints["CreateAccountEndpoint"] = applicationEndpoints.CreateAccountEndpoint(accountService)
	endpoints["GetAccountEndpoint"] = applicationEndpoints.GetAccountEndpoint(accountService)

	return server.NewHTTPServer(ctx, endpoints)
}
