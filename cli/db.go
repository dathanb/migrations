package cli

import (
	"github.com/spf13/cobra"
	"github.com/udacity/migration-demo/db"
	"github.com/udacity/migration-demo/config"
	"fmt"
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
			panic(err.Error())
		}

		err = db.InitDAL(cfg)
		if err != nil {
			panic(err.Error())
		}

		dal := db.ApplicationDAL()

		if dbMigrateUp {
			err = dal.Migrations().MigrateUp()
			if err != nil {
				panic(err.Error())
			}
		} else if dbMigrateDown {
			err = dal.Migrations().MigrateDown()
			if err != nil {
				panic(err.Error())
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
			panic(err.Error())
		}

		err = db.InitDAL(cfg)
		if err != nil {
			panic(err.Error())
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

