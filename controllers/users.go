package controllers

import (
	"MusicIn/models"
	"crypto/sha1"
	"encoding/hex"
	"log"
)

func (c *MainController) Login() {

	c.TplName = "login.html"

}

func (c *MainController) SignUp() {
	c.TplName = "signup.html"
}

func (c *MainController) VerifyUser() {

	email := c.GetString("email")
	password := c.GetString("password")

	h := sha1.New()
	h.Write([]byte(password))
	sha1Hash := hex.EncodeToString(h.Sum(nil))

	id := models.VerifyUser(email, sha1Hash)

	if id != 0 {
		c.TplName = "Home.html"
	}

}

func (c *MainController) CreateUser() {

	newUser := models.UserLogin{}

	// if c.Ctx.Input.Method() == "POST" {
	newUser.Name = c.GetString("name")
	newUser.EmailId = c.GetString("email")

	h := sha1.New()
	h.Write([]byte(c.GetString("password")))
	sha1Hash := hex.EncodeToString(h.Sum(nil))

	newUser.Password = sha1Hash
	newUser.Id = 1

	log.Println(sha1Hash)
	log.Println("Password", newUser.Password)

	models.CreateUser(newUser)

	// }

	c.TplName = "Home.html"

}
