package main

import (
	"github.com/urfave/cli"
	"go-api-demo/internal/config"
	"os"
)

func main()  {
	app := cli.NewApp()
	app.Name = "go-api-demo"
	app.Usage = "go-api-demo -e dev|pro"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name: "env, e",
			Value: "dev",
			Usage: "env: dev|pro",
		},
	}
	app.Action = func(c *cli.Context) error {
		env := c.String("env")
		if err := config.Init(env); err != nil {
			return err
		}

		return nil
	}
	app.Run(os.Args)
}