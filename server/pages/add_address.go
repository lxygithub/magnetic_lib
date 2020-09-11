package pages

import (
	"fmt"
	"github.com/server/models"
	"github.com/server/utils"
	"net/http"
)

func AddAddress(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	if r.Method == "POST" {
		name := r.Form.Get("name")
		url := r.Form.Get("url")
		desc := r.Form.Get("desc")
		address := models.Address{
			Url:  url,
			Name: name,
			Desc: desc,
		}
		if addAddress(address) == nil {
			json, e := utils.Json(address)
			if e == nil {
				fmt.Fprintf(w, json)
			} else {
				fmt.Fprintf(w, utils.ErrJson(500, "错误"))
			}
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
