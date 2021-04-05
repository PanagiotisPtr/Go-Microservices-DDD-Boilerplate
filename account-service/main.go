package main

import (
	"account-service/account"
	"account-service/config"
	"context"
	"database/sql"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func main() {
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewSyncLogger(logger)
		logger = log.With(logger,
			"service", "account",
			"time:", log.DefaultTimestampUTC,
			"called", log.DefaultCaller,
		)
	}
	configFilename := *flag.String("config", "./config/service.dev.yml", "configuration file")
	flag.Parse()

	if os.Getenv("CONFIG_FILE") != "" {
		configFilename = os.Getenv("CONFIG_FILE")
	}

	fmt.Println(configFilename)
	viper.SetConfigFile(configFilename)
	viper.AddConfigPath(".")
	var config config.Configuration

	if err := viper.ReadInConfig(); err != nil {
		level.Error(logger).Log("error", "Failed to load config file for service")
		level.Error(logger).Log("error", err)
		os.Exit(-1)
	}

	if err := viper.Unmarshal(&config); err != nil {
		level.Error(logger).Log("error", "Failed to parse config file for service")
		os.Exit(-1)
	}

	dbsource := "postgresql://" + config.Database.Username + ":" + config.Database.Password +
		"@" + config.Database.Host + ":" + config.Database.Port +
		"/" + config.Database.Name + "?sslmode=disable"

	level.Info(logger).Log("msg", "service started!")
	defer level.Info(logger).Log("msg", "service ended")

	var db *sql.DB
	{
		var err error

		db, err = sql.Open("postgres", dbsource)
		if err != nil {
			level.Error(logger).Log("exit", err)
			os.Exit(-1)
		}
	}

	ctx := context.Background()
	var srv account.Service
	{
		repository := account.NewRepo(db, logger)
		srv = account.NewService(repository, logger)
	}

	errs := make(chan error)

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	endpoints := account.MakeEndpoints(srv)

	go func() {
		fmt.Println("Listening on port", config.Service.Port)
		handler := account.NewHTTPServer(ctx, endpoints)
		errs <- http.ListenAndServe(":"+config.Service.Port, handler)
	}()

	level.Error(logger).Log("exit", <-errs)
}
