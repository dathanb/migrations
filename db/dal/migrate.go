package dal

import (
	"github.com/rubenv/sql-migrate"
	"github.com/jmoiron/sqlx"
	"github.com/dathanb/fakestack/config"
	"github.com/udacity/go-errors"
	"github.com/ansel1/merry"
	"github.com/sirupsen/logrus"
)

const migrationsTable = "migrations"

type MigrationDAL interface {
	MigrateUp() error
	MigrateDown() error
	GetDBVersion() (string, error)
}

type MigrationDALImpl struct {
	db *sqlx.DB
	cfg config.DbConfig
}

func NewMigrationDAL(db *sqlx.DB, cfg config.DbConfig) MigrationDAL {
	return &MigrationDALImpl{db: db, cfg: cfg}
}

func (migrationDAL *MigrationDALImpl) MigrateUp() error {
	migrations := &migrate.FileMigrationSource{
		Dir: "db/migrations",
	}
	migrate.SetTable(migrationsTable)

	count, err := migrate.ExecMax(migrationDAL.db.DB, migrationDAL.cfg.DriverName(), migrations, migrate.Up, 0)
	if err != nil {
		return errors.WithRootCause(merry.New("Failed to migrate"), err)
	}

	logrus.Info("Migrated %d files\n", count)

	return nil
}

func (migrationDAL *MigrationDALImpl) MigrateDown() error {
	migrations := &migrate.FileMigrationSource{
		Dir: "db/migrations",
	}
	migrate.SetTable(migrationsTable)

	count, err := migrate.ExecMax(migrationDAL.db.DB, migrationDAL.cfg.DriverName(), migrations, migrate.Down, 1)
	if err != nil {
		return errors.WithRootCause(merry.New("Failed to migrate"), err)
	}

	logrus.Info("Migrated %d files\n", count)

	return nil
}

func (migrationDAL *MigrationDALImpl) GetDBVersion() (string, error) {
	rows, err := migrationDAL.db.Query("select max(id) from migrations")
	if err != nil {
		return "", err
	}

	rows.Next()
	var version string
	err = rows.Scan(&version)
	if err != nil {
		return "", err
	}

	return version, nil
}
