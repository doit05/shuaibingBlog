package controllers

import (
	"github.com/gin-gonic/gin"
	"shuaibingBlog/helper"
	"shuaibingBlog/helper/apicode"
	"shuaibingBlog/models"
)

func GetUserByName(c *gin.Context) {
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
}
