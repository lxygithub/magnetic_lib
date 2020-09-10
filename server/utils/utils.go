package utils

import (
	"fmt"
	"github.com/spf13/viper"
	"net/http"
	"os"
	"strings"
)

func NeedJson(r http.Request) bool {
	return strings.Contains(r.Header.Get("content-type"), "json")
}

func ReadMySqlConfig(key string) string {
	return ReadConfig(key, "config", "config/tsconfig.json", "json")
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


