package pages

import (
	"github.com/server/utils"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	//t, _ := template.ParseFiles("web/home.html")
	//t.Execute(w,)
}
var engine = utils.GetXORMEngine()
func GetHomeData() {
	defer engine.Close()

	//engine.q

}
