package weixin

import (
	"github.com/silenceper/wechat/menu"
	"shuaibingBlog/utils"
)

func (wxs *WeixinService) UpdateMenu() (err error) {
	mu := utils.Weixin.GetMenu()
	buttons := make([]*menu.Button, 1)
	btn := new(menu.Button)

	//创建click类型菜单
	btn.SetClickButton("name", "key123")
	buttons[0] = btn

	//设置btn为二级菜单
	btn2 := new(menu.Button)
	btn2.SetSubButton("subButton", buttons)

	buttons2 := make([]*menu.Button, 1)
	buttons2[0] = btn2

	//发送请求
	err = mu.SetMenu(buttons2)
	if err != nil {
		utils.Log.Println(err)
	}
	return

}
