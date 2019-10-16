package main

import (
	_ "MusicIn/models"
	_ "MusicIn/routers"

	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}
