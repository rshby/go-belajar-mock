package database

import (
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var DatabaseMysql *gorm.DB

func InitializeMySQLDatabase() (*gorm.DB, func()) {
	db, err := gorm.Open(mysql.Open("root:root@tcp(localhost:3306)/product_db?charset=utf8mb4&parseTime=True&loc=UTC"))
	if err != nil {
		logrus.Fatalf("failed to connect db mysql : %s", err.Error())
	}

	dbMysql, err := db.DB()
	if err != nil {
		logrus.Fatalf("failed to get instance DB : %s", err.Error())
	}

	dbMysql.SetMaxOpenConns(100)
	dbMysql.SetMaxIdleConns(50)
	dbMysql.SetConnMaxIdleTime(30 * time.Minute)
	dbMysql.SetConnMaxLifetime(1 * time.Hour)

	DatabaseMysql = db

	logrus.Info("success connect to MYSQL")
	return DatabaseMysql, func() {
		dbMysql.Close()
	}
}
