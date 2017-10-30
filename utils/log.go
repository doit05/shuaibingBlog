package utils

import (
	"log"
	"os"
	"shuaibingBlog/helper"
)

var Log *log.Logger

func init() {
	file, err := os.Create(helper.Config.LogPath)
	helper.CheckErr(err)
	Log = log.New(file, "shuaibingBlog", log.LstdFlags|log.Llongfile)
}
