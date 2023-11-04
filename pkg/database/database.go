package database

import (
	"app/config"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"net/url"
)

var Orm *gorm.DB

func NewMysqlDB() {
	config.Database.CheckPort()
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=%s",
		config.Database.Username,
		config.Database.Password,
		config.Database.Host,
		config.Database.Port,
		config.Database.Database,
		url.QueryEscape(config.App.Timezone))
	conn, err := gorm.Open(mysql.New(mysql.Config{DSN: dsn}), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	if err != nil {
		panic("error: " + err.Error())
	}
	Orm = conn
}
