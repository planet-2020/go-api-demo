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
	Port string `json:"port"`	//web服务端口
	StatPort string `json:"stat_port"`	//运行统计服务端口
	Env string `json:"env"`	//环境模式
	Debug bool `json:"debug"`	//调试模式
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
	viper.SetDefault("App.Name","go-api-demo")
	viper.SetDefault("App.Port","5050")
	viper.SetDefault("App.StatPort","5051")
	// mysql
	viper.SetDefault("Mysql.Charset","utf8mb4")
	viper.SetDefault("Mysql.MaxIdleConn",5)
	viper.SetDefault("Mysql.MaxOpenConn",10)
	viper.SetDefault("Mysql.ConnMaxLifeTime",600)
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