package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
	"os"
)

var db *gorm.DB

func init() {
	env_err := godotenv.Load()
	if env_err != nil {
		panic("Error in gototenv.Load")
	}

	// "root:111111@tcp(172.17.0.3:3306)/jxc?charset=utf8&parseTime=True&loc=Local"
	dsn := os.Getenv("DB_NAME") + ":" + os.Getenv("DB_PASSWORD") + "@tcp(" +
		os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT") + ")/jxc?charset=utf8&parseTime=True&loc=Local"
	var err interface{}
	db, err = gorm.Open("mysql", dsn)
	if err != nil {
		panic("Error in gorm.Open")
	}

	if os.Getenv("DEBUG") == "true" {
		db.LogMode(true)
	}
}

func GetDB() *gorm.DB {
	return db
}
