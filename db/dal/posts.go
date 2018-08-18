package dal

import (
	"github.com/udacity/migration-demo/models"
	"github.com/jmoiron/sqlx"
	"context"
	"github.com/udacity/go-errors"
	"github.com/ansel1/merry"
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
	tx, err := _dal.db.Begin()
	if err != nil {
		return models.Post{}, errors.WithRootCause(merry.New("failed to begin transaction"), err)
	}

	defer tx.Rollback()

	rows, err := _dal.db.QueryContext(ctx, "select 1 from posts where id = $1", id)
	if err != nil {
		return models.Post{}, errors.WithRootCause(errors.SQLSelectError, err)
	}

	defer rows.Close()
	if rows.Next() {
		_, err = _dal.db.NamedExec(`update posts set post_type = :post_type, user_id = :user_id, 
body = :body where id = :id`, map[string]interface{}{
			"id":           id,
			"post_type":    postType,
			"user_id":      userId,
			"body":         body,
		})
	} else {
		_, err = _dal.db.NamedExec(`insert into users(id, post_type, user_id, body) 
values (:id, :post_type, :user_id, :body)`, map[string]interface{}{
			"id":           id,
			"post_type":    postType,
			"user_id":      userId,
			"body":         body,
		})
	}

	if err != nil {
		return models.Post{}, errors.WithRootCause(merry.New("failed to insert post"), err)
	}

	err = tx.Commit()
	if err != nil {
		return models.Post{}, errors.WithRootCause(merry.New("failed to insert post"), err)
	}

	return models.Post{Id: id, PostType: postType, UserId: userId, Body: body}, nil
}
