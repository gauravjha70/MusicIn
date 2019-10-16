package models

import (
	"log"

	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type UserLogin struct {
	Id       int    `gorm:"primary_key" "auto_increment" "default:1"`
	Name     string `orm:"size(30)"`
	Password string `orm:"size(50)"`
	EmailId  string `orm:"size(50)"`
}

func CreateUser(newUser UserLogin) {

	log.Println("Create User Called")
	db, err := gorm.Open("mysql", beego.AppConfig.String("DbConnection"))

	defer db.Close()
	if err != nil {
		log.Println("Error connecting to Database")
	} else {
		db.Debug().Create(newUser)
	}
}

func VerifyUser(email string, password string) int {

	db, err := gorm.Open("mysql", beego.AppConfig.String("DbConnection"))

	defer db.Close()
	if err != nil {
		log.Println("Error connecting to the Database")
		return 0
	} else {
		login := UserLogin{}
		db.Where("email= ? and password= ?", email, password).First(&login)
		return login.Id
	}

}
