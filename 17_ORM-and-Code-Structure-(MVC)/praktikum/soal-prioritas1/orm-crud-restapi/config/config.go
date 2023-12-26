package config

import (
	"fmt"
	"log"
	"orm-crud-restapi/model"

	// "gorm.io/driver/mysql"
	// "gorm.io/gorm"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	DB *gorm.DB
)

func Init() {

	InitDB()

	InitialMigration()

}

type Config struct {
	DB_Username string

	DB_Password string

	DB_Port string

	DB_Host string

	DB_Name string
}

func InitDB() {

	config := Config{

		DB_Username: "root",

		DB_Password: "ibu123",

		DB_Port: "3306",

		DB_Host: "localhost",

		DB_Name: "crud_go",
	}

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",

		config.DB_Username,

		config.DB_Password,

		config.DB_Host,

		config.DB_Port,

		config.DB_Name,
	)

	var err error

	DB, err = gorm.Open("mysql", connectionString)

	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

}

func InitialMigration() {

	DB.AutoMigrate(&model.User{})
	DB.AutoMigrate(&model.Books{})
	DB.AutoMigrate(&model.Blog{})

}
