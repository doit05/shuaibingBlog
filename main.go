package main

import (
	"github.com/gin-gonic/gin"
	"github.com/robvdl/pongo2gin"
	"shuaibingBlog/helper"
	_ "shuaibingBlog/helper"
	"shuaibingBlog/routers"
)

func main() {
	// Disable Console Color
	// gin.DisableConsoleColor()
	e := gin.Default()
	conf := helper.Config
	e.HTMLRender = pongo2gin.New(pongo2gin.RenderOptions{
		TemplateDir: conf.TemplateDir,
		ContentType: "text/html; charset=utf-8",
	})
	routers.RegisterRouters(e)
	// Listen and Server in 0.0.0.0:8080
	e.Run("0.0.0.0:" + conf.Port)
}