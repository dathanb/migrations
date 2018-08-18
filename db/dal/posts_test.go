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

	t.Run("Inserts new post if one doesn't already exist", func(t *testing.T) {
		mock.ExpectBegin()

		rows := sqlmock.
			NewRows([]string{""})

		mock.ExpectQuery("select 1 from posts where id = \\$1").
			WithArgs(1).
			WillReturnRows(rows) // Empty result set means no matching records

		mock.ExpectExec("insert into posts\\(id, post_type, user_id, body\\) values \\(\\?, \\?, \\?, \\?\\)").
			WithArgs(1, 1, 1, "Test post").
			WillReturnResult(sqlmock.NewResult(1, 1))

		mock.ExpectCommit()

		dal := NewPostsDAL(db)
		_, err := dal.UpsertPost(context.TODO(), 1, 1, 1, "Test post")
		assert.NoError(t, err)

		assert.NoError(t, mock.ExpectationsWereMet(), "Mock Expectations")
	})

	t.Run("Updates existing post", func(t *testing.T) {
		mock.ExpectBegin()

		rows := sqlmock.
			NewRows([]string{""}).
			AddRow(1)

		mock.ExpectQuery("select 1 from posts where id = \\$1").
			WithArgs(1).
			WillReturnRows(rows)

		mock.ExpectExec("update posts set post_type = \\?, user_id = \\?, body = \\? where id = \\?").
			WithArgs(1, 1, "Test post", 1).
			WillReturnResult(sqlmock.NewResult(1, 1))

		mock.ExpectCommit()

		dal := NewPostsDAL(db)
		_, err := dal.UpsertPost(context.TODO(), 1, 1, 1, "Test post")
		assert.NoError(t, err)

		assert.NoError(t, mock.ExpectationsWereMet(), "Mock Expectations")
	})
}
