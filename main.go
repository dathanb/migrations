package main

import (
	"github.com/udacity/migration-demo/cli"
	"os"
	"github.com/sirupsen/logrus"
)

func main() {
	if len(os.Args) == 1 {
		os.Args = append(os.Args, "help")
	}

	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetLevel(logrus.DebugLevel)

	cli.Run(os.Args[1:])
}
