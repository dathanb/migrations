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

	t.Run("Inserts new user if one doesn't already exist", func(t *testing.T) {
		mock.ExpectBegin()

		rows := sqlmock.
			NewRows([]string{""})

		mock.ExpectQuery("select 1 from users where id = \\$1").
			WithArgs(1).
			WillReturnRows(rows) // Empty result set means no matching records

		mock.ExpectExec("insert into users\\(id, display_name\\) values \\(\\?, \\?\\)").
			WithArgs(1, "Test user").
			WillReturnResult(sqlmock.NewResult(1, 1))

		mock.ExpectCommit()

		dal := NewUsersDAL(db)
		dal.UpsertUser(context.TODO(), 1, "Test user")

		assert.NoError(t, mock.ExpectationsWereMet(), "Mock Expectations")
	})

	t.Run("Updates existing user", func(t *testing.T) {
		mock.ExpectBegin()

		rows := sqlmock.
			NewRows([]string{""}).
			AddRow(1)

		mock.ExpectQuery("select 1 from users where id = \\?").
			WithArgs(1).
			WillReturnRows(rows)

		mock.ExpectExec("insert into users\\(id, display_name\\) values \\(\\?, \\?\\)").
			WithArgs(1, "Test user").
			WillReturnResult(sqlmock.NewResult(1, 1))

		mock.ExpectCommit()

		dal := NewUsersDAL(db)
		dal.UpsertUser(context.TODO(), 1, "Test user")

		assert.NoError(t, mock.ExpectationsWereMet(), "Mock Expectations")
	})
}
