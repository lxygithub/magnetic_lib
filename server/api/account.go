package api

import (
	"fmt"
	"github.com/server/models"
	"github.com/server/utils"
	"net/http"
)

func UserManagerRegister(w http.ResponseWriter, r *http.Request) {

}

func ManagerLogin(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	if r.Method == "POST" {
		username := r.Form.Get("username")
		password := r.Form.Get("password")
		err, user := FindManager(username, utils.Md5Salt(password))
		token, _ := utils.CreateToken(
			user["id"],
			username,
		)
		if err == nil {
			result, _ := utils.Json(models.BaseResp{
				Msg:   "登陆成功",
				Data:  user,
				Token: string(utils.AesEncryptCFB([]byte(token))),
			})
			w.Header().Set("content-type", "application/json; charset=UTF-8")
			fmt.Fprintf(w, result)
		}
	}

}

func FindManager(username, password string) (error, map[string]string) {
	engine := utils.GetXORMEngine()
	result, err := engine.QueryString(fmt.Sprintf(`select uid,name,username,enable,remove,verify,root,create_time,updatetime 
												from user where username="%s" and password="%s"`, username, password))
	if err == nil && len(result) > 0 {
		return err, result[0]
	}
	return err, map[string]string{}
}
func UserLogin(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

}

func AddManager(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

}
