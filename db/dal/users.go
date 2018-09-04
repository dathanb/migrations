package dal

import (
	"context"
	"github.com/ansel1/merry"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
	"github.com/udacity/go-errors"
	"github.com/udacity/migration-demo/models"
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
		if log.GetLevel() >= log.ErrorLevel {
			log.Errorf("Could not begin transaction to upsert user %d", id)
		}
		return models.User{}, errors.WithRootCause(merry.New("failed to begin transaction"), err)
	}

	defer tx.Rollback()

	if log.GetLevel() >= log.DebugLevel {
		log.WithField("id", id).WithField("displayName", displayName).Debug("Inserting user")
	}
	_, err = _dal.db.NamedExec("insert into users(id, display_name) values (:id, :display_name) on conflict (id) do update set display_name = EXCLUDED.display_name", map[string]interface{}{
		"id":           id,
		"display_name": displayName,
	})

	if err != nil {
		if log.GetLevel() >= log.ErrorLevel {
			log.Errorf("Got error while upserting user with id %d: %s", id, err.Error())
		}
		return models.User{}, errors.WithRootCause(merry.New("failed to insert user"), err)
	}

	err = tx.Commit()
	if err != nil {
		return models.User{}, errors.WithRootCause(merry.New("failed to insert user"), err)
	}

	return models.User{Id: id, DisplayName: displayName}, nil
}
