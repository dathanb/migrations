package cli

import (
	"github.com/spf13/cobra"
	"github.com/dathanb/migrations/fakestack/api"
	"github.com/dathanb/migrations/fakestack/config"
	"github.com/dathanb/migrations/fakestack/db"
	"github.com/sirupsen/logrus"
)

var startCmd = &cobra.Command{
	Use: "start",
	Short: "run server",
	Long: "Run the service and block",
	Run: func(cmd *cobra.Command, args []string) {
		logrus.SetFormatter(&logrus.JSONFormatter{})

		cfg, err := config.LoadConfig()
		if err != nil {
			panic(err.Error())
		}

		err = db.InitDAL(cfg)
		if err != nil {
			panic(err)
		}



		api.Serve(cfg.Server.Port())
	},
}
