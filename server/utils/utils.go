package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/server/models"
	"github.com/spf13/viper"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
)

const AES_KEY = "PEIVOAJRTEAOGHOA"

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

func JsonString(model interface{}) string {
	bytes, _ := json.Marshal(model)
	return string(bytes)
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

func ToInt64(str string) int64 {
	i, _ := strconv.ParseInt(str, 10, 64)
	return i
}
func ToInt(str string) int {
	i, _ := strconv.Atoi(str)
	return i
}

func AesEncryptCFB(origData []byte) []byte {
	block, err := aes.NewCipher([]byte(AES_KEY))
	if err != nil {
		panic(err)
	}
	encrypted := make([]byte, aes.BlockSize+len(origData))
	iv := encrypted[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(encrypted[aes.BlockSize:], origData)
	return encrypted
}
func AesDecryptCFB(encrypted []byte) (decrypted []byte, err error) {
	block, _ := aes.NewCipher([]byte(AES_KEY))
	if len(encrypted) < aes.BlockSize {
		return nil, errors.New("ciphertext too short")
	}
	iv := encrypted[:aes.BlockSize]
	encrypted = encrypted[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(encrypted, encrypted)
	return encrypted, nil
}

/**
hex解码和aes解密
*/
func DecryptHexAesStr(hexAesTokenStr string) (tokenStr string, err error) {
	aesTokenStr, err := hex.DecodeString(hexAesTokenStr)
	bytes, err := AesDecryptCFB(aesTokenStr)
	if err == nil {
		tokenStr = string(bytes)
	}
	return
}
