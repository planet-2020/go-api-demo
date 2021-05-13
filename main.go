package main

import (
	"github.com/urfave/cli"
	"go-api-demo/internal/config"
	"go-api-demo/models"
	"go-api-demo/router"
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
		// 初始化配置文件
		if err := config.Init(env); err != nil {
			return err
		}

		//初始化数据库
		if err := models.Database(config.Conf.Mysql); err != nil {
			return err
		}

		//初始化路由并启动web服务
		server := router.Init(config.Conf.App)
		server.GinEngine.Run(":"+config.Conf.App.Port)

		return nil
	}
	app.Run(os.Args)
}