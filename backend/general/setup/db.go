package setup

import (
	"general/constants"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitGorm() *gorm.DB {
	var db *gorm.DB
	var err error
	for {
		dsn := "root:" + constants.MYSQL_PASSWORD + "@(" + constants.MYSQL_IP + ")/go_project?charset=utf8&parseTime=True&loc=Local"
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Println(err)
			time.Sleep(5 * time.Second)
			continue
		}
		return db
	}
}
