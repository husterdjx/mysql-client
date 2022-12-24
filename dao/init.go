package dao

import (
	"fmt"
	"myclient/conf"
	"myclient/model"
	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
	RS *RdbService
)

type RdbService struct {
	tx *gorm.DB
}

func init() {
	connStr := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true&charset=utf8mb4,utf8",
		conf.C.DB.UserName,
		conf.C.DB.Password,
		conf.C.DB.Host,
		conf.C.DB.Schema,
	)
	engine, err := gorm.Open(mysql.Open(connStr), &gorm.Config{})
	if err != nil {
		logrus.Fatalf("[init]database connect error %v", err)
		panic(err)
	}
	DB = engine
	RS = &RdbService{
		DB,
	}
	sqldb, err := engine.DB()
	if err != nil {
		logrus.Fatalf("[init]invalid database driver %v", err)
		panic(err)
	}
	sqldb.SetMaxIdleConns(10)
	sqldb.SetMaxOpenConns(10000)
	sqldb.SetConnMaxLifetime(time.Second * 3)
	DB.AutoMigrate(model.Course{})
	DB.AutoMigrate(model.SC{})
	DB.AutoMigrate(model.Student{})

	// logrus.Info("[init] db init")
}
