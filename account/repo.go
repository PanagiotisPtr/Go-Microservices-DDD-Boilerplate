package account

import (
	"context"
	"database/sql"
	"errors"

	"github.com/go-kit/kit/log"
)

var RepoErr = errors.New("Unable to handle repo request")

type repo struct {
	db     *sql.DB
	logger log.Logger
}

func NewRepo(db *sql.DB, logger log.Logger) Repository {
	return &repo{
		db:     db,
		logger: log.With(logger, "repo", "sql"),
	}
}

func (repo *repo) GetUser(ctx context.Context, id string) (string, error) {
	var email string
	err := repo.db.QueryRow(`SELECT email from users where id=$1`, id).Scan(&email)

	if err != nil {
		return "", RepoErr
	}

	return email, nil
}

func (repo *repo) CreateUser(ctx context.Context, user User) error {
	sql := `insert into users (id, email, password) values ($1, $2, $3)`

	repo.logger.Log("QUERY: ", sql)

	if user.Email == "" || user.Password == "" {
		return RepoErr
	}

	_, err := repo.db.ExecContext(ctx, sql, user.ID, user.Email, user.Password)

	if err != nil {
		return err
	}

	return nil
}
