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
	DB       *gorm.DB
	RS       *RdbService
	username string
	password string
)

type RdbService struct {
	tx *gorm.DB
}

func init() {
	ui()
	login()
	connStr := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true&charset=utf8mb4,utf8",
		username,
		password,
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

/*打印一个ASCII art图像，内容为"husterdjx"*/
func ui() {
	fmt.Println("  __  __           _       _   _      _   _             _ ")
	fmt.Println(" |  \\/  |         | |     | | | |    | | | |           | |")
	fmt.Println(" | \\  / | ___   __| | __ _| |_| | ___| |_| |__   ___   | |")
	fmt.Println(" | |\\/| |/ _ \\ / _` |/ _` | __| |/ _ \\ __| '_ \\ / _ \\  | |")
	fmt.Println(" | |  | | (_) | (_| | (_| | |_| |  __/ |_| | | |  __/  |_|")
	fmt.Println(" |_|  |_|\\___/ \\__,_|\\__,_|\\__|_|\\___|\\__|_| |_|\\___|  (_)")
	fmt.Println("                                                          ")
}

/*实现一个登录界面*/
func login() {
	fmt.Println("请输入数据库用户名：")
	fmt.Scanln(&username)
	fmt.Println("请输入密码：")
	fmt.Scanln(&password)
}
