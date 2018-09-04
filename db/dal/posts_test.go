package dal

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"context"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestPostgresPostsDAL_UpsertPost(t *testing.T) {
	var err error
	db, mock, err := getMockDB()
	assert.NoError(t, err, "Failed to mock the db")

	t.Run("Upserts posert", func(t *testing.T) {
		mock.ExpectBegin()

		mock.ExpectExec("insert into posts\\(id, post_type, user_id, body\\) values \\(\\?, \\?, \\?, \\?\\) on conflict \\(id\\) do update set post_type = EXCLUDED.post_type, user_id = EXCLUDED.user_id, body = EXCLUDED.body").
			WithArgs(1, 1, 1, "Test post").
			WillReturnResult(sqlmock.NewResult(1, 1))

		mock.ExpectCommit()

		dal := NewPostsDAL(db)
		_, err := dal.UpsertPost(context.TODO(), 1, 1, 1, "Test post")
		assert.NoError(t, err)

		assert.NoError(t, mock.ExpectationsWereMet(), "Mock Expectations")
	})
}
