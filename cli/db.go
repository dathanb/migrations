package cli

import (
	"github.com/spf13/cobra"
	"github.com/dathanb/fakestack/db"
	"github.com/dathanb/fakestack/config"
	"fmt"
	"github.com/udacity/go-errors"
	"github.com/ansel1/merry"
)

var dbCmd = &cobra.Command{
	Use: "db",
	Short: "run database commands",
	Long: "Run database commands (e.g., migrations)",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Usage()
	},
}

var (
	dbMigrateUp bool
	dbMigrateDown bool
)

var dbMigrateCmd = &cobra.Command{
	Use: "migrate",
	Short: "Run database migrations",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := config.LoadConfig()
		if err != nil {
			panic(errors.WithRootCause(merry.New("Failed to load config"), err))
		}

		err = db.InitDAL(cfg)
		if err != nil {
			panic(errors.WithRootCause(merry.New("Failed to initialize data access layer"), err))
		}

		dal := db.ApplicationDAL()

		if dbMigrateUp {
			err = dal.Migrations().MigrateUp()
			if err != nil {
				panic(errors.WithRootCause(merry.New("Failed to migrate"), err))
			}
		} else if dbMigrateDown {
			err = dal.Migrations().MigrateDown()
			if err != nil {
				panic(errors.WithRootCause(merry.New("Failed to migrate"), err))
			}
		} else {
			fmt.Print("Must specify one of --up or --down")
			cmd.Usage()
			return
		}
	},
}

var dbVersionCmd = &cobra.Command{
	Use: "version",
	Short: "show database version",
	Run: func(cmd *cobra.Command, args[] string) {
		cfg, err := config.LoadConfig()
		if err != nil {
			panic(errors.WithRootCause(merry.New("Failed to load config"), err))
		}

		err = db.InitDAL(cfg)
		if err != nil {
			panic(errors.WithRootCause(merry.New("Failed to initialize data access layer"), err))
		}

		version, err := db.ApplicationDAL().Migrations().GetDBVersion()
		if err != nil {
			panic(err.Error())
		}

		fmt.Printf("DB version: %s\n", version)
	},
}

func init() {
	dbMigrateCmd.Flags().BoolVarP(&dbMigrateUp, "up", "u", false, "run upward migrations")
	dbMigrateCmd.Flags().BoolVarP(&dbMigrateDown, "down", "d", false, "rollback one migration")
	dbCmd.AddCommand(dbMigrateCmd)

	dbCmd.AddCommand(dbVersionCmd)
}

