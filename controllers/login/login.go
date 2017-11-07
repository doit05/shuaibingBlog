package login

import (
	"github.com/flosch/pongo2"
	"github.com/gin-gonic/gin"
	"shuaibingBlog/helper/apicode"
)

func Login(c *gin.Context) {
	data := pongo2.Context{}
	data["site"] = c.Request.Host
	c.HTML(apicode.Success, "login.html", data)

}
