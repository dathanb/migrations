package db

import (
	_ "github.com/lib/pq"
	"github.com/jmoiron/sqlx"
	"github.com/udacity/migration-demo/db/dal"
	"github.com/udacity/migration-demo/config"
)

var globalDal DAL

type DAL interface {
	Migrations() dal.MigrationDAL
	Users() dal.UsersDAL
}

type dalImpl struct {
	db *sqlx.DB
	cfg *config.Config
}

func connect(driverName string, connectionString string) (*sqlx.DB, error) {
	return sqlx.Connect(driverName, connectionString)
}

// Returns the process-global DAL instance
func ApplicationDAL() DAL {
	return globalDal
}

func (appDAL *dalImpl) Migrations() dal.MigrationDAL {
	return dal.NewMigrationDAL(appDAL.db, appDAL.cfg.Db)
}

func (appDAL *dalImpl) Users() dal.UsersDAL {
	return dal.NewUsersDAL(appDAL.db)
}

func InitDAL(cfg *config.Config) error {
	db, err := connect(cfg.Db.DriverName(), cfg.Db.ConnectionString())
	if err != nil {
		return err
	}

	globalDal = &dalImpl{db, cfg}

	return nil
}

func init() {
	globalDal = nil
}
