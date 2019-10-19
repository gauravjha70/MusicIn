package controllers

import (
	"MusicIn/models"
	"crypto/sha1"
	"encoding/hex"
	"time"
)

func (c *MainController) Login() {
	sess := c.GetSession("musicInLogin")
	if sess == nil {
		c.TplName = "login.html"
	} else {
		c.Redirect("/home", 303)
	}

	var dates [31]int
	for i := 0; i < 31; i++ {
		dates[i] = i + 1
	}
	c.Data["dates"] = dates

	var years [100]int
	j := 1920
	for i := 0; i < 100; i++ {
		years[i] = j
		j = j + 1
	}
	c.Data["years"] = years

}

func (c *MainController) SignUp() {
	c.TplName = "signup.html"
}

//Verify a user using login credentials and start the session
func (c *MainController) VerifyUser() {

	email := c.GetString("email")
	password := c.GetString("password")

	h := sha1.New()
	h.Write([]byte(password))
	sha1Hash := hex.EncodeToString(h.Sum(nil))

	id := models.VerifyUser(email, sha1Hash)

	if id != 0 {
		//Setting up the session
		userSessions := models.UserSession{}
		userSessions.Email = email
		userSessions.Id = id
		userSessions.TimeStamp = time.Now().String()
		c.SetSession("musicInLogin", userSessions)
		c.Redirect("/home", 303)
	}

}

//Create a new user
func (c *MainController) CreateUser() {

	newUser := models.UserLogin{}
	userDetails := models.UserDetails{}
	if c.Ctx.Input.Method() == "POST" {
		newUser.EmailId = c.GetString("useremail")

		h := sha1.New()
		h.Write([]byte(c.GetString("userpass")))
		sha1Hash := hex.EncodeToString(h.Sum(nil))

		newUser.Password = sha1Hash
		newUser.Id = 1

		userDetails.FirstName = c.GetString("userfirstname")
		userDetails.LastName = c.GetString("userlastname")
		userDetails.NickName = c.GetString("usernickname")
		day := c.GetString("selectday")
		month := c.GetString("selectmonth")
		year := c.GetString("selectyear")
		userDetails.Dob = day + "/" + month + "/" + year
		userDetails.Gender = c.GetString("usergender")
		userDetails.Genre = c.GetString("userPrefferedGenre")
		userDetails.About = c.GetString("userabout")
		userDetails.Status = c.GetString("userstatus")
		models.CreateUser(newUser, userDetails)

	}

	c.TplName = "Home.html"

}

//Logouts the user
func (c *MainController) Logout() {
	c.DelSession("musicInLogin")
	c.Redirect("/", 303)
}
