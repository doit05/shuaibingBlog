package weixin

import (
	"github.com/silenceper/wechat/message"
)

type WeixinService struct {
}

func (wxs *WeixinService) MessageHandler(msg message.MixMessage) (ret *message.Reply) {
	switch msg.MsgType {
	//文本消息
	case message.MsgTypeText:
		//回复消息：演示回复用户发送的消息
		text := message.NewText(msg.Content)
		ret = &message.Reply{MsgType: message.MsgTypeText, MsgData: text}

		//图片消息
	case message.MsgTypeImage:
		//mediaID 可通过素材管理-上上传多媒体文件获得
		image := message.NewImage("mediaID")
		return &message.Reply{message.MsgTypeVideo, image}

		//语音消息
	case message.MsgTypeVoice:
		//do something
		music := message.NewMusic("title", "description", "musicURL", "hQMusicURL", "thumbMediaID")
		return &message.Reply{message.MsgTypeMusic, music}

		//视频消息
	case message.MsgTypeVideo:
		//do something
		video := message.NewVideo("mediaID", "视频标题", "视频描述")
		return &message.Reply{message.MsgTypeVideo, video}

		//小视频消息
	case message.MsgTypeShortVideo:
		//do something

		//地理位置消息
	case message.MsgTypeLocation:
		//do something

		//链接消息
	case message.MsgTypeLink:
		//do something

		//事件推送消息
	case message.MsgTypeEvent:

	}

	return

}
