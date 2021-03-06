package api

import (
	"fmt"
	"github.com/server/models"
	"github.com/server/utils"
	"html/template"
	"net/http"
	"time"
)

func AddAddress(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	if r.Method == "GET" {
		t, _ := template.ParseFiles("web/add_address.html")
		t.Execute(w, "")
	} else if r.Method == "POST" {
		name := r.PostForm.Get("name")
		url := r.PostForm.Get("url")
		desc := r.PostForm.Get("desc")
		address := models.Address{
			Url:     url,
			Name:    name,
			Desc:    desc,
			Enable:  true,
			Created: time.Now().Unix(),
		}
		if addAddress(address) == nil {
			json := utils.JsonString(models.BaseResp{
				Code:   0,
				ErrMsg: "",
				Msg:    "",
				Data:   address,
			})
			fmt.Fprintf(w, json)

		}
	}
}

func addAddress(address models.Address) error {
	engine := utils.GetXORMEngine()
	defer engine.Close()

	_, err := engine.InsertOne(address)
	if err != nil {
		return err
	}
	return nil

}
