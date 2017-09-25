package login

import (
	"github.com/gin-gonic/gin"
	"shuaibingBlog/helper/apicode"
)

func Login(c *gin.Context) {
	c.String(200, "pong")

}


