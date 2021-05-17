package models

import (
	"fmt"
	"go-api-demo/internal/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var DB *gorm.DB

/**
 * @Description: 连接mysql
 * @param mysqlConfig
 * @return error
 */
func Database(mysqlConfig config.MysqlConfig) error {
	// 连接mysql
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s",
		mysqlConfig.Username,
		mysqlConfig.Password,
		mysqlConfig.Hostname,
		mysqlConfig.Port,
		mysqlConfig.Database,
		mysqlConfig.Charset)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	// 设置连接池
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}
	// 空闲连接池
	sqlDB.SetMaxIdleConns(mysqlConfig.MaxIdleConn)
	// 最大连接池
	sqlDB.SetMaxOpenConns(mysqlConfig.MaxOpenConn)
	// 超时时间
	sqlDB.SetConnMaxLifetime(time.Second * time.Duration(mysqlConfig.ConnMaxLifeTime))

	DB = db
	return nil
}
