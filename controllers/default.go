package controllers

import (
	"github.com/gin-gonic/gin"
	//"mysite/helper/apicode"
	//"shuaibingBlog/helper"
	//"shuaibingBlog/models"
	"github.com/flosch/pongo2"
)

//// Get user value
func Index(c *gin.Context) {
	c.HTML(200, "index.html", pongo2.Context{"name": "world"})

}

func Ping(c *gin.Context) {
	c.String(200, "pong")
}
