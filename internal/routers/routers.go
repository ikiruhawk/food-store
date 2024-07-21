package routers

import (
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
)

var tpl *template.Template

func RunRouters() *mux.Router {
	tpl, _ = template.ParseGlob("views/html/*.html")
	r := mux.NewRouter()
	r.HandleFunc("/", homeHandleFunc).Methods("GET")
	r.HandleFunc("/products", productsHandleFunc).Methods("GET")
	return r
}

func homeHandleFunc(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	tpl.ExecuteTemplate(w, "index.html", nil)
}

func productsHandleFunc(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	tpl.ExecuteTemplate(w, "products.html", nil)
}
