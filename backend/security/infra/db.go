package infra

import (
	"log"
	"security/constants"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var _ RDB = (*MySQL)(nil)

type MySQL struct {
	db *gorm.DB
}

func (m *MySQL) Run() (err error) {
	for {
		dsn := "root:" + constants.MYSQL_PASSWORD + "@(" + constants.MYSQL_IP + ")/go_project?charset=utf8&parseTime=True&loc=Local"
		m.db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Println(err)
			time.Sleep(5 * time.Second)
			continue
		}
		break
	}
	return
}

func (m *MySQL) Stop() {
	if m.db == nil {
		return
	}

	db, err := m.db.DB()
	if err == nil {
		db.Close()
	}
}

func (m *MySQL) Orm() *gorm.DB {
	return m.db
}
