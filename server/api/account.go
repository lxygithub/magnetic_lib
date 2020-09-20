package api

import (
	"encoding/hex"
	"fmt"
	"github.com/server/models"
	"github.com/server/utils"
	"html/template"
	"net/http"
)

func Register(w http.ResponseWriter, r *http.Request) {

}

func Login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	if r.Method == "POST" {
		username := r.Form.Get("username")
		password := r.Form.Get("password")
		err, user := FindUser(username, utils.Md5Salt(password))

		if err == nil {
			token, _ := utils.CreateToken(
				models.User{
					Uid:      user["uid"],
					Name:     user["name"],
					Username: user["username"],
					Enable:   utils.ToInt(user["enable"]),
					Remove:   utils.ToInt(user["remove"]),
					Verify:   utils.ToInt(user["verify"]),
					Root:     utils.ToInt(user["root"]),
				},
			)
			result := utils.JsonString(models.BaseResp{
				Msg:   "登陆成功",
				Data:  user,
				Token: hex.EncodeToString(utils.AesEncryptCFB([]byte(token))),
			})
			w.Header().Set("content-type", "application/json; charset=UTF-8")
			fmt.Fprintf(w, result)
		}
	} else {
		t, _ := template.ParseFiles("web/login.html")
		t.Execute(w, nil)
	}

}

func FindUser(username, password string) (error, map[string]string) {
	engine := utils.GetXORMEngine()
	result, err := engine.QueryString(fmt.Sprintf(`select uid,name,username,enable,remove,verify,root,create_time,updatetime 
												from user where username="%s" and password="%s"`, username, password))
	if err == nil && len(result) > 0 {
		return err, result[0]
	}
	return err, map[string]string{}
}

func AddUser(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

}
