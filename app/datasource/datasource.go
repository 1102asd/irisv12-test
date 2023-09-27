package datasource

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"irisv12-test/app/conf"
	"strings"
	"time"
)

var db *gorm.DB

func GetDB() *gorm.DB {
	return db
}

func init() {
	path := strings.Join([]string{conf.MysqlConfig.UserName, ":", conf.MysqlConfig.Password, "@(", conf.MysqlConfig.Hostname, ":", conf.MysqlConfig.Port, ")/", conf.MysqlConfig.Database, "?charset=utf8&parseTime=true"}, "")
	var err error
	db, err = gorm.Open("mysql", path)
	if err != nil {
		fmt.Println("dsadsadsadsad", conf.MysqlConfig)
		fmt.Println("dsadsadsadsad")
		fmt.Println("dsadsadsadsad")
		panic(err)
	}
	db.SingularTable(true)
	db.DB().SetConnMaxLifetime(1 * time.Second)
	db.DB().SetMaxIdleConns(100)  //最大打开的连接数
	db.DB().SetMaxOpenConns(2000) //设置最大闲置个数
	db.SingularTable(true)        //表生成结尾不带s
	// 启用Logger，显示详细日志
	db.LogMode(true)
}
