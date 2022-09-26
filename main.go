package main

import (
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"rest-skeleton/api"
)

func main() {
	app := &cli.App{
		Flags: []cli.Flag{
			&cli.UintFlag{
				Name:  "port",
				Usage: "application port",
				Value: 8080,
			},
		},
		Action: func(cCtx *cli.Context) error {
			return api.Startup(cCtx)
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
