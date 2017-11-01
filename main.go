package main

import (
	"github.com/gin-gonic/gin"
	"github.com/robvdl/pongo2gin"
	"shuaibingBlog/helper"
	"shuaibingBlog/routers"
	_ "shuaibingBlog/utils"
)

func main() {
	e := gin.Default()

	//获取配置
	conf := helper.Config
	gin.SetMode(conf.Mod)
	if conf.Mod == gin.ReleaseMode {
		gin.DisableConsoleColor() // Disable Console Color
	}

	//设置模板引擎
	e.HTMLRender = pongo2gin.New(pongo2gin.RenderOptions{
		TemplateDir: conf.TemplateDir,
		ContentType: "text/html; charset=utf-8",
	})

	//注册路由
	routers.RegisterRouters(e)
	// Listen and Server in 0.0.0.0:8080
	e.Run("0.0.0.0:" + conf.Port)
}
