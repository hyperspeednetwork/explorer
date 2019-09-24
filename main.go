package main

import (
	"github.com/wongyinlong/hsnNet/crawler"
	_ "github.com/wongyinlong/hsnNet/routers"

	"github.com/astaxie/beego"
)

// entry ,start beego and crawler
func main() {
	crawler.OnStart()
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}

