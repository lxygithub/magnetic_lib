package api

import (
	"fmt"
	"github.com/server/models"
	"github.com/server/utils"
	"html/template"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	if r.Method == "GET" {
		t, _ := template.ParseFiles("web/home.html")
		t.Execute(w, "")
	} else {
		addresses, err := GetHomeData()
		if err == nil {
			data := utils.JsonString(models.BaseResp{
				Code:   0,
				ErrMsg: "",
				Data:   addresses,
			})
			fmt.Fprintf(w, data)
		}
	}

}
func GetHomeData() ([]models.Address, error) {
	var engine = utils.GetXORMEngine()
	defer engine.Close()

	var addresses []models.Address
	err := engine.Find(&addresses)
	return addresses, err

}
