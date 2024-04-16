package database

import (
	"fmt"
	"lms_backend/global"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var DBconn *gorm.DB

func InitPostgres() {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s", global.Conf.DBHost, global.Conf.DBPort, global.Conf.DBUserName, global.Conf.DBUserPassword, global.Conf.DBName)

	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   fmt.Sprintf("%s.%s_", global.Conf.DBSchema, global.Conf.DBTablePrefix),
			SingularTable: true,
		},
	})

	if err != nil {
		panic(err)
	}
	log.Println("Postgres Database Connected")
	conn.Logger = logger.Default.LogMode(logger.Info)

	sqlDB, _ := conn.DB()

	if val := global.Conf.DBMaxIdleConn; val != 0 {
		sqlDB.SetMaxIdleConns(val)
	}

	if val := global.Conf.DBMaxOpenConn; val != 0 {
		sqlDB.SetMaxOpenConns(val)
	}

	if val := global.Conf.DBMaxConnLifetime; val != 0 {
		sqlDB.SetConnMaxLifetime(time.Duration(val) * time.Second)
	}

	if val := global.Conf.DBMaxOpenConn; val != 0 {
		sqlDB.SetMaxOpenConns(val)
	}

	DBconn = conn
}
