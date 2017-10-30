package helper

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"regexp"
	"strings"
	"time"
)

func GetType(url string) (goaltype int64) {
	goaltype = 1
	re := regexp.MustCompile("monitor")
	if re.MatchString(url) {
		goaltype = 2
	}
	return goaltype
}

// md5加密
func Md5(s string) string {
	return Md5Bytes([]byte(s))
}

// md5加密
func Md5Bytes(s []byte) string {
	h := md5.New()
	h.Write(s)
	return hex.EncodeToString(h.Sum(nil))
}

// 判断某个字符串是否在slice中, 类似php的in_array()函数
func InSlice(value string, slice []string) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

// 判断某个对象是否在slice中, 类似php的in_array()函数
func InSliceObj(value interface{}, slice []interface{}) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

// 写入log到文件
func Log2File(file string, content string) {

	// log 目录
	logDir := GetRootPath() + "/logs/" + time.Now().Format("20060102")

	// 创建目录
	err := os.MkdirAll(logDir, 0666)
	if err != nil {
		log.Fatalf("error can't mkdir logs: %v", err)
	}

	// log文件路径
	file = logDir + "/" + file + ".log"

	// 打开文件
	f, ferr := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if ferr != nil {
		log.Fatalf("error opening file: %v", ferr)
	}
	defer f.Close()

	// 记录错误到文件
	log.SetOutput(f)

	log.Println(" " + content)
}

// 获取入口文件的绝对路径
func GetRootPath() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return dir
}

// 是否是生产环境
func IsProductionEnv() bool {
	if os.Getenv("PAY_HOST") == "prod" {
		// 生产环境
		return true
	}
	// 开发或测试环境
	return false
}

// 获取配置的前缀
// 例如开发环境返回 dev:: 生产环境返回 prod::
func GetConfigPrifix() string {
	str := "::"
	if IsProductionEnv() {
		str = "prod" + str
	} else {
		str = "dev" + str
	}

	return str
}

// 手机号码格式是否正确
func IsMobile(mobile string) bool {
	var validID = regexp.MustCompile(`^\d{11}$`)
	return validID.MatchString(mobile)
}

type ApiRes struct {
	Code string      `json:"status"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type ApiResArr struct {
	Code string        `json:"status"`
	Msg  string        `json:"msg"`
	Data []interface{} `json:"data"`
}

// 初始化api返回数据
func InitApiRes(code, msg string, data interface{}) *ApiRes {
	apiRes := new(ApiRes)
	apiRes.Code = code
	apiRes.Msg = msg
	apiRes.Data = data
	return apiRes
}

func GetFloat(unk interface{}) (float64, error) {
	var floatType = reflect.TypeOf(float64(0))
	v := reflect.ValueOf(unk)
	v = reflect.Indirect(v)
	if !v.Type().ConvertibleTo(floatType) {
		return 0, fmt.Errorf("cannot convert %v to float64", v.Type())
	}
	fv := v.Convert(floatType)
	return fv.Float(), nil
}

func checkFileExist(path string) (ret bool, err error) {
	_, err = os.Open(path) // For read access.
	ret = err == nil
	return

}

// fileName:文件名字(带全路径)
// content: 写入的内容
func appendToFile(fileName string, content string) error {
	// 以只写的模式，打开文件
	f, err := os.OpenFile(fileName, os.O_WRONLY, 0644)
	if err != nil {
		return err
	} else {
		// 查找文件末尾的偏移量
		n, _ := f.Seek(0, os.SEEK_END)
		// 从末尾的偏移量开始写入内容
		_, err = f.WriteAt([]byte(content), n)
	}
	defer f.Close()
	return err
}

//读取文件需要经常进行错误检查，这个帮助方法可以精简下面的错误检查过程。
func CheckErr(e error) {
	if e != nil {
		panic(e)
	}
}

func GetCurrentPath() string {
	s, err := exec.LookPath(os.Args[0])
	CheckErr(err)
	i := strings.LastIndex(s, "\\")
	path := string(s[0 : i+1])
	return path
}
