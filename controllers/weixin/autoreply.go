package weixin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"shuaibingBlog/helper/apicode"
	"shuaibingBlog/service/weixin"
	"shuaibingBlog/utils"
)

func Weixin(c *gin.Context) {

	wc := utils.Weixin
	wxs := &weixin.WeixinService{}

	// 传入request和responseWriter
	server := wc.GetServer(c.Request, c.Writer)
	//设置接收消息的处理方法
	server.SetMessageHandler(wxs.MessageHandler)

	//处理消息接收以及回复
	err := server.Serve()
	if err != nil {
		fmt.Println(err)
		return
	}
	//发送回复的消息
	server.Send()
}

func GetAccessToken(c *gin.Context) {
	wc := utils.Weixin
	accessToken, err := wc.GetAccessToken()
	if err != nil {
		utils.Log.Println(err)
		c.String(apicode.QueryError, err.Error())
	}
	c.String(200, accessToken)
}
