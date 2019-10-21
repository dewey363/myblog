package main

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/plugins/cors"
	"github.com/sinksmell/LanBlog/models"
)

func init() {
	models.RegisterDB()
	orm.Debug = false
	orm.RunSyncdb("default", false, true)
}

func main() {
	fmt.Println("hello myblog!")
	//设置访问过滤
	beego.InsertFilter("*",beego.BeforeRouter,cors.Allow(
		&cors.Options{
			AllowAllOrigins:true,
			AllowMethods: []string{"GET","POST","OPTIONS"},
			AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
			ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
			AllowCredentials: true,
		}))

	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	beego.Run()
}