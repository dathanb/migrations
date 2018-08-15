package dal

import (
	"github.com/udacity/migration-demo/models"
	"github.com/jmoiron/sqlx"
	"context"
	"github.com/udacity/go-errors"
	"github.com/ansel1/merry"
	"fmt"
)

type UsersDAL interface {
	CreateUser(ctx context.Context, id int, display_name string) (models.User, error)
}

func NewUsersDAL(db *sqlx.DB) UsersDAL {
	return &PostgresUsersDAL{
		db: db,
	}
}

type PostgresUsersDAL struct {
	db *sqlx.DB
}

func (_dal *PostgresUsersDAL) CreateUser(ctx context.Context, id int, display_name string) (models.User, error) {
	var err error
	tx, err := _dal.db.Begin()
	if err != nil {
		return nil, errors.WithRootCause(merry.New("failed to begin transaction"), err)
	}

	defer tx.Rollback()

	_, err = _dal.db.NamedExec("insert into users(id, display_name) values (:id, :display_name)", map[string]interface{}{
		"id": id,
		"display_name": display_name,
	})

	if err != nil {
		fmt.Printf("%s\n", err.Error())
		return nil, errors.WithRootCause(merry.New("failed to insert user"), err)
	}

	err = tx.Commit()
	if err != nil {
		return nil, errors.WithRootCause(merry.New("failed to insert user"), err)
	}

	return models.NewUser(id, display_name), nil
}