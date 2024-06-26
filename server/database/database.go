package database

import (
	"log"
	"os"

	"github.com/blalyasar/go-react-blog/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DBConn *gorm.DB

func ConnectDB() {
	// data source name
	user := os.Getenv("db_user")
	dbpass := os.Getenv("db_password")
	dbname := os.Getenv("db_name")

	dsn := user + ":" + dbpass + "@tcp(127.0.0.1:3306)/" + dbname + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Error),
	})
	if err != nil {
		panic("Database connection failed")
	}
	log.Println("Connection succesfull")

	db.AutoMigrate(new(model.Blog))

	DBConn = db

}
