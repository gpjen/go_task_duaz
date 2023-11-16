package db

import (
	"fmt"
	"user-product-management/app/products"
	"user-product-management/app/users"
	"user-product-management/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitMysqlDB() {

	env := config.LoadEnvConfig()

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", env.DBUsername, env.DBPassword, env.DBHost, env.DBPort, env.DBDatabase)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("success to connect database")

	DB = db

	MigrateModels()
}

func MigrateModels() {
	DB.AutoMigrate(&users.User{}, &products.Product{})
}
