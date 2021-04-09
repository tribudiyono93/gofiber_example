package database

import (
	"fmt"
	"github.com/tribudiyono93/gofiber_example/fiber-rest-api/entity"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *gorm.DB

const (
	maxIdleCons = 10
	maxOpenCons = 10
)

func Connect()  {
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbName)
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to create a connection to database")
	}

	DB.AutoMigrate(&entity.User{}, &entity.UserModuleRole{}, &entity.Book{})

	//set connection pooling
	sql, err := DB.DB()
	if err != nil {
		panic("failed to get sql db")
	}
	sql.SetMaxIdleConns(maxIdleCons)
	sql.SetMaxOpenConns(maxOpenCons)
	log.Println("successfully create database connection")
}

func Close() {
	sql, err := DB.DB()
	if err != nil {
		panic("failed to get sql from DB()")
	}
	if err := sql.Close(); err != nil {
		panic("failed to close database connection")
	}
	log.Println("successfully close database connection")
}
