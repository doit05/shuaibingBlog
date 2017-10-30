package main

import (
	"github.com/gin-gonic/gin"
	"github.com/robvdl/pongo2gin"
	"shuaibingBlog/routers"
)

func main() {
	// Disable Console Color
	// gin.DisableConsoleColor()
	e := gin.Default()
	e.HTMLRender = pongo2gin.New(pongo2gin.RenderOptions{
		TemplateDir: "shuaibingBlog/views",
		ContentType: "text/html; charset=utf-8",
	})
	routers.RegisterRouters(e)
	// Listen and Server in 0.0.0.0:8080
	e.Run(":8080")
}
