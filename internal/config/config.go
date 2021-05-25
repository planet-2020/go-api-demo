package config

import (
	"errors"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"strings"
)

const (
	confPath = "./config"
)

var (
	Conf *Config
)
var Envs = map[string]string{
	"local":"config.local",
	"dev":"config.dev",
	"prod":"config.prod",
}

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
	JwtSubject string `json:"jwt_subject"`	//jwt主题
	JwtSecret string `json:"jwt_secret"`	//jwt密钥
	JwtExpireTime int64 `json:"jwt_expire_time"`	//jwt过期时间，秒
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
	viper.SetDefault("App.JwtSubject","go-api-demo")
	viper.SetDefault("App.JwtSecret","go123456")
	viper.SetDefault("App.JwtExpireTime","86400")
	// mysql
	viper.SetDefault("Mysql.Charset","utf8mb4")
	viper.SetDefault("Mysql.MaxIdleConn",5)
	viper.SetDefault("Mysql.MaxOpenConn",10)
	viper.SetDefault("Mysql.ConnMaxLifeTime",600)
}

/**
 * @Description: 配置文件初始化
 * @param env 环境模式
 * @return error
 */
func Init(env string) error {
	if env == "" {
		env = "local"
	} else {
		env = strings.ToLower(env)
	}
	confFile, ok := Envs[env]
	if !ok {
		return errors.New("环境变量错误")
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