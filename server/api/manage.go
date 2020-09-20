package api

import (
	"fmt"
	"github.com/server/models"
	"github.com/server/utils"
	"html/template"
	"net/http"
	"strings"
)

func Manage(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	if r.Method == "POST" {
		hexAesToken := r.Form.Get("token")
		if hexAesToken != "" {
			token, err := utils.DecryptHexAesStr(hexAesToken)
			if err == nil {
				c, _ := utils.ParseToken(token)
				root := utils.GetTokenVal(c, "root")
				enable := utils.GetTokenVal(c, "enable")
				remove := utils.GetTokenVal(c, "remove")
				if root == 1 || (enable == 1 && remove == 1) {
					t, _ := template.ParseFiles("web/manage.html")
					t.Execute(w, nil)
				}
			}

		}
	}

}

func Delete(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	if r.Method == "POST" {
		//err := json.Unmarshal([]byte(r.PostFormValue("ids")), &ids)
		token, err := utils.DecryptHexAesStr(r.Form.Get("token"))
		if err != nil {
			fmt.Fprintf(w, utils.JsonString(models.BaseResp{Code: 403, ErrMsg: err.Error()}))
			return
		}
		c, _ := utils.ParseToken(token)
		if utils.GetTokenVal(c, "remove") == 1 {
			ids := r.Form.Get("ids")
			if ids != "" {
				ids = strings.ReplaceAll(ids, "[", "")
				ids = strings.ReplaceAll(ids, "]", "")
				engine := utils.GetXORMEngine()
				defer engine.Close()
				sql := fmt.Sprintf("DELETE FROM cili_engine WHERE id IN (%s);", ids)
				_, err := engine.Exec(sql)
				if err == nil {
					resp := utils.JsonString(models.BaseResp{
						Msg: "删除成功",
					})
					fmt.Fprintf(w, resp)
				}
			}
		} else {
			fmt.Fprintf(w, utils.JsonString(models.BaseResp{
				Code:   403,
				ErrMsg: "没有删除权限",
			}))
		}
	}
}
