package main

import (
	//"fmt"
	"github.com/gin-gonic/gin"
	"mysite/helper/apicode"
	"shuaibingBlog/helper"
	"shuaibingBlog/models"
)

func main() {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	//// Get user value
	r.GET("/user/:name", func(c *gin.Context) {
		nickname := c.Params.ByName("name")
		model := new(models.UserModel)
		user := models.User{}
		user.Nickname = nickname
		ok := model.FindByName(&user)
		var ret *helper.ApiRes
		if ok == nil && user.Id > 0 {
			ret = helper.InitApiRes(apicode.Success, apicode.Msg(apicode.Success), user)
		} else {
			ret = helper.InitApiRes(apicode.QueryError, apicode.Msg(apicode.QueryError), ok)
		}
		c.JSON(200, *ret)

	})

	// Authorized group (uses gin.BasicAuth() middleware)
	// Same than:
	// authorized := r.Group("/")
	// authorized.Use(gin.BasicAuth(gin.Credentials{
	//	  "foo":  "bar",
	//	  "manu": "123",
	//}))
	authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
		"foo":  "bar", // user:foo password:bar
		"manu": "123", // user:manu password:123
	}))

	authorized.POST("admin", func(c *gin.Context) {
		c.MustGet(gin.AuthUserKey)
		// Parse JSON
		var json struct {
			Value string `json:"value" binding:"required"`
		}

		if c.Bind(&json) == nil {
			c.JSON(200, gin.H{"status": json.Value})
		}

	})

	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
