package database

import (
	"estate/options"
	"log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDb() (*gorm.DB, error) {
	dsn := options.DBuser + ":" + options.DBpass + "@tcp(" + options.DBhost + ")/" + options.DBname + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return db, err
}

func CloseDb(db *gorm.DB) {
	sqlDB, err := db.DB()
	if err != nil {
		log.Println("Database shutdown error:", err)
	}
	defer sqlDB.Close()
}
