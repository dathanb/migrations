package main

import (
	"github.com/udacity/migration-demo/cli"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		os.Args = append(os.Args, "help")
	}

	//cfg, _ := config.LoadConfig()

	cli.Run(os.Args[1:])

	//connStr := fmt.Sprintf("user=%s password=%s sslmode=disable", config.Db.Username, config.Db.Password)
	//db, err := sql.Open("postgres", connStr)
	//bailOnError(err)
	//defer db.Close()

	//ctx := context.Background()
	//err = saveGithubData(ctx, config, db)

}
