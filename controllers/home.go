package controllers

import "MusicIn/models"

func (c *MainController) HomePage() {

	sess := c.GetSession("musicInLogin")
	if sess == nil {
		c.Redirect("/", 303)
	}
	c.TplName = "Home.html"
	userSession := sess.(models.UserSession)
	c.Data["UserName"] = userSession.Email
	// c.Data["AllNames"] = models.GetAllUsers()
	var allNames [3]string
	allNames[0] = "Hello"
	allNames[1] = "World"
	allNames[2] = "Hi"

	c.Data["AllNames"] = allNames

}
