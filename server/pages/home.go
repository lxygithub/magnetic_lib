package pages

import (
	"html/template"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	t, _ := template.ParseFiles("web/home.html")
	t.Execute(w, "")
}
func GetHomeData() {
	//var engine = utils.GetXORMEngine()
	//defer engine.Close()

	//engine.q

}
