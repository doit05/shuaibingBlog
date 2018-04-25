package weixin

import (
	"github.com/gin-gonic/gin"
	"shuaibingBlog/helper/apicode"
	"shuaibingBlog/service/weixin"
)

func UpdateMenu(c *gin.Context) {
	wxs := &weixin.WeixinService{}
	err := wxs.UpdateMenu()
	if err != nil {
		c.JSON(apicode.QueryError, err.Error())
	}
	c.JSON(apicode.Success, apicode.Msg(apicode.Success))

}
