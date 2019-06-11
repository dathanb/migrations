package cli

import (
	"fmt"
	"github.com/ansel1/merry"
	"github.com/dathanb/migrations/fakestack/config"
	"github.com/dathanb/migrations/fakestack/db"
	"github.com/spf13/cobra"
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
			panic(merry.WithUserMessage(err, "Failed to load config"))
		}

		err = db.InitDAL(cfg)
		if err != nil {
			panic(merry.WithMessage(err, "Failed to initialize data access layer"))
		}

		dal := db.ApplicationDAL()

		if dbMigrateUp {
			err = dal.Migrations().MigrateUp()
			if err != nil {
				panic(merry.WithMessage(err, "Failed to migrate"))
			}
		} else if dbMigrateDown {
			err = dal.Migrations().MigrateDown()
			if err != nil {
				panic(merry.WithMessage(err, "Failed to migrate"))
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
			panic(merry.WithUserMessage(err, "Failed to load config"))
		}

		err = db.InitDAL(cfg)
		if err != nil {
			panic(merry.WithMessage(err, "Failed to initialize data access layer"))
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

