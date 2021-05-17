package main

import (
	"fmt"
	"github.com/arl/statsviz"
	"github.com/urfave/cli"
	"go-api-demo/internal/config"
	"go-api-demo/models"
	"go-api-demo/router"
	"net/http"
	"os"
)

func main()  {
	app := cli.NewApp()
	app.Name = "go-api-demo"
	app.Usage = "go-api-demo -e local|dev|prod"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name: "env, e",
			Value: "local",
			Usage: "env: local|dev|prod",
		},
	}
	app.Action = func(c *cli.Context) error {
		env := c.String("env")
		// 初始化配置文件
		if err := config.Init(env); err != nil {
			fmt.Println(err)
			return err
		}

		//初始化数据库
		if err := models.Database(config.Conf.Mysql); err != nil {
			fmt.Println(err)
			return err
		}

		//初始化路由并启动web服务
		server := router.Init(config.Conf.App)
		if config.Conf.App.Debug == true {
			go runtimeStatistics(config.Conf.App.StatPort)
		}
		server.GinEngine.Run(":"+config.Conf.App.Port)

		return nil
	}
	app.Run(os.Args)
}

/**
 * @Description: 运行实时统计
 * @param port 端口
 */
func runtimeStatistics(port string)  {
	//运行后访问 http://127.0.0.1:5051/debug/statsviz/ 查看统计数据
	statsviz.RegisterDefault()
	http.ListenAndServe(":"+port,nil)
}