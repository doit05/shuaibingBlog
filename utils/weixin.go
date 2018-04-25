package utils

import (
	"github.com/silenceper/wechat"
	"shuaibingBlog/helper"
)

var Weixin *wechat.Wechat

func init() {
	//配置微信参数
	config := &wechat.Config{
		AppID:          helper.Config.Wechatconf.AppID,
		AppSecret:      helper.Config.Wechatconf.AppSecret,
		Token:          helper.Config.Wechatconf.Token,
		EncodingAESKey: helper.Config.Wechatconf.EncodingAESKey,
	}

	config.Cache = Cache

	wc := wechat.NewWechat(config)
	if wc == nil {
		panic("init wechat error")
	}
	Weixin = wc
}
