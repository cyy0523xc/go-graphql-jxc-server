package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
	"os"

	"app/database"
)

func main() {
	env_err := godotenv.Load()
	if env_err != nil {
		println("Error in gototenv.Load")
		return
	}

	// "root:111111@tcp(172.17.0.3:3306)/jxc?charset=utf8&parseTime=True&loc=Local"
	dsn := os.Getenv("DB_NAME") + ":" + os.Getenv("DB_PASSWORD") + "@tcp(" +
		os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT") + ")/jxc?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		println("Error in gorm.Open")
		return
	}

	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&database.User{}, &database.Goods{})
}
