package dal

import (
	"github.com/jmoiron/sqlx"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

// Creates a mock db and wraps it in a sqlx connection
func getMockDB() (*sqlx.DB, sqlmock.Sqlmock, error) {
	// Setup the mock db and wrap it with sqlx
	db, mock, err := sqlmock.New()
	return sqlx.NewDb(db, "sqlmock"), mock, err
}


