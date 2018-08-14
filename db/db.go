package db

import (
	_ "github.com/lib/pq"
	"github.com/jmoiron/sqlx"
	"github.com/udacity/migration-demo/db/dal"
	"github.com/udacity/migration-demo/config"
)

type DAL interface {
	GetMigrationDAL() dal.MigrationDAL
}

type dalImpl struct {
	db *sqlx.DB
	cfg *config.Config
}

func connect(driverName string, connectionString string) (*sqlx.DB, error) {
	return sqlx.Connect(driverName, connectionString)
}

func GetDAL(cfg *config.Config) (DAL, error) {
	db, err := connect(cfg.Db.DriverName(), cfg.Db.ConnectionString())
	if err != nil {
		return nil, err
	}

	return &dalImpl{db, cfg}, nil
}

func (appDAL *dalImpl) GetMigrationDAL() dal.MigrationDAL {
	return dal.NewMigrationDAL(appDAL.db, appDAL.cfg.Db)
}

