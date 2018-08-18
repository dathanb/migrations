package dal

import (
	"github.com/udacity/migration-demo/models"
	"github.com/jmoiron/sqlx"
	"context"
	"github.com/udacity/go-errors"
	"github.com/ansel1/merry"
)

type UsersDAL interface {
	UpsertUser(ctx context.Context, id int, displayName string) (models.User, error)
}

func NewUsersDAL(db *sqlx.DB) UsersDAL {
	return &PostgresUsersDAL{
		db: db,
	}
}

type PostgresUsersDAL struct {
	db *sqlx.DB
}

func (_dal *PostgresUsersDAL) UpsertUser(ctx context.Context, id int, displayName string) (models.User, error) {
	var err error
	tx, err := _dal.db.Begin()
	if err != nil {
		return models.User{}, errors.WithRootCause(merry.New("failed to begin transaction"), err)
	}

	defer tx.Rollback()

	rows, err := _dal.db.QueryContext(ctx, "select 1 from users where id = $1", id)
	if err != nil {
		return models.User{}, errors.WithRootCause(errors.SQLSelectError, err)
	}

	defer rows.Close()
	if rows.Next() {
		_, err = _dal.db.NamedExec("update users set display_name = :display_name where id = :id", map[string]interface{}{
			"id":           id,
			"display_name": displayName,
		})
	} else {
		_, err = _dal.db.NamedExec("insert into users(id, display_name) values (:id, :display_name)", map[string]interface{}{
			"id":           id,
			"display_name": displayName,
		})
	}

	if err != nil {
		return models.User{}, errors.WithRootCause(merry.New("failed to insert user"), err)
	}

	err = tx.Commit()
	if err != nil {
		return models.User{}, errors.WithRootCause(merry.New("failed to insert user"), err)
	}

	return models.User{Id: id, DisplayName: displayName}, nil
}
