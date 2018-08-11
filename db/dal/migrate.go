package dal

import (
	migrate "github.com/rubenv/sql-migrate"
	"github.com/jmoiron/sqlx"
	"github.com/udacity/migration-demo/config"
	"github.com/udacity/migration-demo/error"
	"fmt"
)

const migrationsTable = "migrations"

type MigrationDAL interface {
	MigrateUp() error.Error
	MigrateDown() error.Error
	GetDBVersion() (string, error.Error)
}

type MigrationDALImpl struct {
	db *sqlx.DB
	cfg config.DbConfig
}

func NewMigrationDAL(db *sqlx.DB, cfg config.DbConfig) MigrationDAL {
	return &MigrationDALImpl{db: db, cfg: cfg}
}

func (migrationDAL *MigrationDALImpl) MigrateUp() error.Error{
	migrations := &migrate.FileMigrationSource{
		Dir: "db/migrations",
	}
	migrate.SetTable(migrationsTable)

	count, err := migrate.ExecMax(migrationDAL.db.DB, migrationDAL.cfg.DriverName(), migrations, migrate.Up, 0)
	if err != nil {
		return error.WithMessage("Failed to migrate")
	}

	fmt.Printf("Migrated %d files\n", count)

	return nil
}

func (migrationDAL *MigrationDALImpl) MigrateDown() error.Error {
	migrations := &migrate.FileMigrationSource{
		Dir: "db/migrations",
	}
	migrate.SetTable(migrationsTable)

	count, err := migrate.ExecMax(migrationDAL.db.DB, migrationDAL.cfg.DriverName(), migrations, migrate.Down, 1)
	if err != nil {
		return error.WithMessage("Failed to migrate")
	}

	fmt.Printf("Migrated %d files\n", count)

	return nil
}

func (migrationDAL *MigrationDALImpl) GetDBVersion() (string, error.Error) {
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
