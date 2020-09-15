package utils

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/server/models"
	"github.com/spf13/viper"
	"net/http"
	"os"
	"strings"
)

func NeedJson(r *http.Request) bool {
	return strings.Contains(r.Header.Get("content-type"), "json")
}

func ReadMySqlConfig(key string) string {
	return ReadConfig(key, "config", "config", "json")
}

func ReadConfig(key, configName, configPath, configType string) string {
	viper.SetConfigName(configName) //设置配置文件的名字
	viper.AddConfigPath(configPath) //添加配置文件所在的路径
	viper.SetConfigType(configType) //设置配置文件类型，可选
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("config file error: %s\n", err)
		os.Exit(1)
	}
	return viper.GetString(key)
}

func Json(model interface{}) (string, error) {
	bytes, err := json.Marshal(models.BaseResp{
		Code:   0,
		ErrMsg: "",
		Data:   model,
	})
	return string(bytes), err
}
func ErrJson(code int, errMsg string) string {
	errBytes, _ := json.Marshal(models.BaseResp{
		Code:   code,
		ErrMsg: errMsg,
		Data:   nil,
	})
	return string(errBytes)
}

func Md5Salt(str string) string {
	salt := "cili@@lib"
	h := md5.New()
	h.Write([]byte(salt + str))
	return hex.EncodeToString(h.Sum(nil))
}
func Md5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}
