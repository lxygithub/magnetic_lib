package pages

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

	t, _ := template.ParseFiles("/web/manage.html")
	t.Execute(w, nil)

}

func Delete(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	if r.Method == "POST" {
		//err := json.Unmarshal([]byte(r.PostFormValue("ids")), &ids)
		ids := r.Form.Get("ids")
		if ids != "" {
			ids = strings.ReplaceAll(ids, "[", "(")
			ids = strings.ReplaceAll(ids, "]", ")")
			engine := utils.GetXORMEngine()
			defer engine.Close()
			sql := fmt.Sprintf("DELETE FROM cili_engine WHERE id IN %s;", ids)
			_, err := engine.Exec(sql)
			if err == nil {
				resp, _ := utils.Json(models.BaseResp{
					Msg: "删除成功",
				})
				fmt.Fprintf(w, resp)
			}
		}
	}
}
