package apicode

const (
	Success                  int = 0000 // 成功
	SystemError              int = 1    // 系统错误
	MissParam                int = 2    // 参数缺失
	InvalidParam             int = 3    // 参数错误
	ApiNotFound              int = 4    // 接口未找到
	SignError                int = 5    // 签名错误
	RequestExpire            int = 6    // 请求已失效
	QueryError               int = 7    //查询出错
	DeleteError              int = 8    //查询出错
	InsertDataFailed         int = 9    //插入数据失败
	SaveFileFailed           int = 10   //保存文件失败
	DockerClientError        int = 11   //docker client err
	DockerContainerListError int = 12   //docker ContainerList err

	InsertGoalFailed int = 0100 // 插入用户失败
	ReceivedMsg      int = 0101 //消息已经上传
)

// 获取错误代码对应的提示信息
func Msg(code string) string {
	msgMap := map[string]string{
		Success:                  "success",
		SystemError:              "系统错误",
		MissParam:                "参数缺失",
		InvalidParam:             "参数错误",
		ApiNotFound:              "接口未找到",
		SignError:                "签名错误",
		RequestExpire:            "请求已失效",
		QueryError:               "查询出错",
		DeleteError:              "删除出错",
		ReceivedMsg:              "消息已经上传",
		InsertGoalFailed:         "插入用户目标失败",
		InsertDataFailed:         "插入数据失败",
		SaveFileFailed:           "保存文件失败",
		DockerClientError:        "docker client err",
		DockerContainerListError: "docker ContainerList err",
	}

	msg, exist := msgMap[code]

	if exist == false { // 不存在默认返回空字符串
		msg = ""
	}

	return msg
}
