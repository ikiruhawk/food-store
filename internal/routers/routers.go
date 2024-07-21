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
	r.HandleFunc("/article/{id:[0-9]+}/{name}", articleHandleFunc).Methods("GET")
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

type article struct {
	Id   string
	Name string
}

func articleHandleFunc(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	mv := mux.Vars(r)
	a := article{
		Id:   mv["id"],
		Name: mv["name"],
	}
	tpl.ExecuteTemplate(w, "article.html", a)
}
