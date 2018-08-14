package dal

import (
	"testing"
	"os"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"context"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestMain(m *testing.M) {
	log.SetLevel(log.DebugLevel)
	os.Exit(m.Run())
}

func TestCreateUserInsertsUser(t *testing.T) {
	var err error
	db, mock, err := getMockDB()
	assert.NoError(t, err, "Failed to mock the db")

	mock.ExpectBegin()

	mock.ExpectExec("insert into users\\(id, display_name\\) values \\(\\?, \\?\\)").
		WithArgs(1, "Test user").
		WillReturnResult(sqlmock.NewResult(1, 1))


	mock.ExpectCommit()
	mock.ExpectClose()

	dal := NewUsersDAL(db)
	dal.CreateUser(context.TODO(), 1, "Test user")

}

