package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"pandora/conf"
)

var DB *gorm.DB

func init() {
	var err error
	connStr := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		conf.AppConfig.Database.User,
		conf.AppConfig.Database.PassWord,
		conf.AppConfig.Database.Host,
		conf.AppConfig.Database.Port,
		conf.AppConfig.Database.DataBase)
	fmt.Println(connStr)
	DB, err = gorm.Open("mysql", connStr)
	if err != nil {
		panic(err)
	}
	DB.DB().SetMaxIdleConns(10)
	DB.DB().SetMaxOpenConns(100)
}
