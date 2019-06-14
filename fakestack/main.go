package main

import (
	"github.com/sirupsen/logrus"
	"github.com/dathanb/migrations/fakestack/cli"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		os.Args = append(os.Args, "help")
	}

	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetLevel(logrus.DebugLevel)

	cli.Run(os.Args[1:])
}
