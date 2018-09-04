package dal

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"context"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)


func TestPostgresUsersDAL_UpsertUser(t *testing.T) {
	var err error
	db, mock, err := getMockDB()
	assert.NoError(t, err, "Failed to mock the db")

	t.Run("Upserts new user", func(t *testing.T) {
		mock.ExpectBegin()

		mock.ExpectExec("insert into users\\(id, display_name\\) values \\(\\?, \\?\\) on conflict \\(id\\) do update set display_name = EXCLUDED.display_name").
			WithArgs(1, "Test user").
			WillReturnResult(sqlmock.NewResult(1, 1))

		mock.ExpectCommit()

		dal := NewUsersDAL(db)
		_, err := dal.UpsertUser(context.TODO(), 1, "Test user")
		assert.NoError(t, err)

		assert.NoError(t, mock.ExpectationsWereMet(), "Mock Expectations")
	})
}
