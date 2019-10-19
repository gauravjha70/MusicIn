package models

import (
	"log"

	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type UserSession struct {
	Id        int
	Email     string
	TimeStamp string
}

func init() {

	Db, Err := gorm.Open("mysql", beego.AppConfig.String("DbConnection"))

	defer Db.Close()
	if Err != nil {
		log.Println("Connection to DB failed..!")
	}

	log.Println("Connection to DB established....")

	Db.AutoMigrate(&UserLogin{})
	Db.AutoMigrate(&UserDetails{})
	log.Println("Synced with DB.")

}
