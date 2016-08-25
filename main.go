package main

import (
	_ "github.com/zhangjunfang/swagger-go/docs"
	_ "github.com/zhangjunfang/swagger-go/routers"

	"github.com/astaxie/beego"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
