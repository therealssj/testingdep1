package main

import (
	"fmt"
	"os"

	"github.com/therealssj/testingdep1/src/cli"
)

func main() {
	app, err := cli.NewApp()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
