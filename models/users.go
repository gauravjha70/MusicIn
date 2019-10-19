package models

import (
	"log"

	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type UserLogin struct {
	Id       int    `gorm:"primary_key";"AUTO_INCREMENT"`
	Password string `orm:"size(50)"`
	EmailId  string `orm:"size(50)"`
}

type UserDetails struct {
	Id        int    `gorm:"primary_key" "auto_increment"`
	UserId    int    `orm:"size(50)"`
	FirstName string `orm:"size(50)"`
	LastName  string `orm:"size(50)"`
	NickName  string `orm:"size(50)"`
	Dob       string `orm:"size(50)"`
	Gender    string `orm:"size(50)"`
	Genre     string `orm:"size(50)"`
	Status    string `orm:"size(50)"`
	About     string `orm:"size(200)"`
}

func CreateUser(newUser UserLogin, userDetails UserDetails) {

	db, err := gorm.Open("mysql", beego.AppConfig.String("DbConnection"))

	defer db.Close()
	if err != nil {
		log.Println("Error connecting to Database")
	} else {
		temp := UserLogin{}
		temp.Id = 0 //Initialised ID
		tempDetails := UserDetails{}
		tempDetails.Id = 0
		db.Debug().Last(&temp)
		newUser.Id = temp.Id + 1 //New ID Generated
		db.Debug().Create(newUser)
		db.Debug().Last(&tempDetails)
		userDetails.Id = tempDetails.Id + 1 //New ID Generated
		userDetails.UserId = newUser.Id
		db.Debug().Create(userDetails)
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
		db.Debug().Where("email_id= ? and password= ?", email, password).First(&login)
		return login.Id
	}
}

//Get all Users
func GetAllUsers() UserLogin {
	db, err := gorm.Open("mysql", beego.AppConfig.String("DbConnection"))
	users := UserLogin{}
	defer db.Close()
	if err != nil {
		log.Println("Error connecting to the database")
	} else {
		db.Debug().Find(&users)
	}
	return users
}
