package config

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

const (
	confPath = "./config"
)

var (
	Conf *Config
)

type Config struct {
	App struct{
		Name string `json:"name"`
		Debug bool `json:"debug"`
	} `json:"app"`

	Mysql struct{
		Hostname string	`json:"hostname"`
		Database string	`json:"database"`
		Username string	`json:"username"`
		Password string	`json:"password"`
		Port int `json:"port"`
		Prefix string `json:"prefix"`
	} `json:"mysql"`

	Redis struct{
		Host string `json:"host"`
		Port string `json:"port"`
		Auth string `json:"auth"`
		SelectDb int `json:"select_db"`
	} `json:"redis"`
	
}

func setDefault()  {
	viper.SetDefault("app.name","go-api-demo")
}

func Init(env string) error {
	// 获取配置文件
	var confFile string
	switch env {
	case "dev":
		confFile = "config.dev"
	case "pro":
		confFile = "config.pro"
	default:
		confFile = "config.dev"
	}
	viper.AddConfigPath(confPath)
	viper.SetConfigName(confFile)
	viper.SetConfigType("json")
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	setDefault()	// 设置默认值
	if err := viper.Unmarshal(&Conf); err != nil {
		return err
	}
	// 监控重新读取配置文件
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		if err := viper.Unmarshal(&Conf); err != nil {
			panic(err)
		}
	})

	return nil
}