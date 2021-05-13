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
	App AppConfig `json:"app"`
	Mysql MysqlConfig `json:"mysql"`
	Redis RedisConfig `json:"redis"`
}

type AppConfig struct {
	Name string `json:"name"`
	Port string `json:"port"`
	Env string `json:"env"`
	Debug bool `json:"debug"`
}

type MysqlConfig struct {
	Hostname string	`json:"hostname"`
	Database string	`json:"database"`
	Username string	`json:"username"`
	Password string	`json:"password"`
	Port int `json:"port"`
	Charset string `json:"charset"`
	Prefix string `json:"prefix"`
	MaxIdleConn int `json:"max_idle_conn"`
	MaxOpenConn int `json:"max_open_conn"`
	ConnMaxLifeTime int `json:"conn_max_life_time"`
}

type RedisConfig struct {
	Host string `json:"host"`
	Port string `json:"port"`
	Auth string `json:"auth"`
	SelectDb int `json:"select_db"`
}

func setDefault()  {
	// app
	viper.SetDefault("app.name","go-api-demo")
	viper.SetDefault("app.port","5050")
	// mysql
	viper.SetDefault("mysql.charset","utf8mb4")
	viper.SetDefault("mysql.max_idle_conn",5)
	viper.SetDefault("mysql.max_open_conn",10)
	viper.SetDefault("mysql.conn_max_life_time",600)
}

func Init(env string) error {
	// 获取配置文件
	var confFile string
	switch env {
	case "dev":
		confFile = "config.dev"
		viper.SetDefault("app.env","dev")
	case "pro":
		confFile = "config.pro"
		viper.SetDefault("app.env","pro")
	default:
		confFile = "config.dev"
		viper.SetDefault("app.env","dev")
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