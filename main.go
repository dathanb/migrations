package main

import (
	"os"

)

func main() {
	if len(os.Args) == 1 {
		os.Args = append(os.Args, "help")
	}

	//config := LoadConfig()
	//fmt.Println(config)

	//connStr := fmt.Sprintf("user=%s password=%s sslmode=disable", config.Db.Username, config.Db.Password)
	//db, err := sql.Open("postgres", connStr)
	//bailOnError(err)
	//defer db.Close()

	//ctx := context.Background()
	//err = saveGithubData(ctx, config, db)

}
