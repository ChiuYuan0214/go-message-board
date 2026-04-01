package infra

import (
	"general/constants"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var _ RDB = (*MySQL)(nil)

type MySQL struct {
	db *gorm.DB
}

func (i *MySQL) Run() (err error) {
	for {
		dsn := "root:" + constants.MYSQL_PASSWORD + "@(" + constants.MYSQL_IP + ")/go_project?charset=utf8&parseTime=True&loc=Local"
		i.db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Println(err)
			time.Sleep(5 * time.Second)
			continue
		}
		break
	}

	return
}

func (i *MySQL) Stop() {
	db, err := i.db.DB()
	if err != nil {
		db.Close()
	}
}

func (i *MySQL) Orm() *gorm.DB {
	return i.db
}
