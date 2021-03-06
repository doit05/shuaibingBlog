package helper

import (
	"github.com/BurntSushi/toml"
	"io/ioutil"
	"os"
)

type Conf struct {
	ID          int          `toml:"ID"`
	Mod         string       `toml:"Mod"`
	Port        string       `toml:"Port"`
	AppName     string       `toml:"AppName"`
	TemplateDir string       `toml:"TemplateDir"`
	LogPath     string       `toml:"LogPath"`
	Mysqlcon    string       `toml:"Mysqlcon"`
	Redisconn   string       `toml:"Redisconn"`
	Wechatconf  WechatConfig `toml:"wechat"`
}

type WechatConfig struct {
	AppID          string
	AppSecret      string
	Token          string
	EncodingAESKey string
}

var Config *Conf

func init() {
	ReadConf()
}

func ReadConf() (c *Conf, err error) {
	fname := GetCurrentPath() + "conf/conf.toml"
	var (
		fp       *os.File
		fcontent []byte
	)
	c = new(Conf) // &Conf{}
	if fp, err = os.Open(fname); err != nil {
		return
	}

	if fcontent, err = ioutil.ReadAll(fp); err != nil {
		return
	}

	if err = toml.Unmarshal(fcontent, c); err != nil {
		return
	}
	Config = c
	return
}
