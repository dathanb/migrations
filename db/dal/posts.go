package dal

import (
	"context"
	"github.com/ansel1/merry"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
	"github.com/udacity/go-errors"
	"github.com/dathanb/fakestack/models"
)

type PostsDAL interface {
	UpsertPost(ctx context.Context, id int, postType int, userId int, body string) (models.Post, error)
}

func NewPostsDAL(db *sqlx.DB) PostsDAL {
	return &PostgresPostsDAL{
		db: db,
	}
}

type PostgresPostsDAL struct {
	db *sqlx.DB
}

func (_dal *PostgresPostsDAL) UpsertPost(ctx context.Context, id int, postType int, userId int, body string) (models.Post, error) {
	var err error
	tx, err := _dal.db.Beginx()
	if err != nil {
		if log.GetLevel() >= log.ErrorLevel {
			log.Errorf("Could not begin transaction to upsert post with id %d", id)
		}
		return models.Post{}, errors.WithRootCause(merry.New("Failed to begin transaction"), err)
	}

	defer tx.Rollback()

	if log.GetLevel() >= log.DebugLevel {
		log.WithField("id", id).WithField("post_type", postType).
			WithField("user_id", userId).WithField("body", body).
			Debug("Inserting user")
	}
	_, err = _dal.db.NamedExec(`insert into posts(id, post_type, user_id, body) 
values (:id, :post_type, :user_id, :body)
on conflict (id) do update set
post_type = EXCLUDED.post_type,
user_id = EXCLUDED.user_id,
body = EXCLUDED.body`, map[string]interface{}{
		"id":           id,
		"post_type":    postType,
		"user_id":      userId,
		"body":         body,
	})

	if err != nil {
		if log.GetLevel() >= log.ErrorLevel {
			log.Errorf("Got error while upserting post with id %d: %s", id, err.Error())
		}
		return models.Post{}, errors.WithRootCause(merry.New("failed to insert post"), err)
	}

	err = tx.Commit()
	if err != nil {
		return models.Post{}, errors.WithRootCause(merry.New("failed to insert post"), err)
	}

	return models.Post{Id: id, PostType: postType, UserId: userId, Body: body}, nil
}
