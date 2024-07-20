package routers

import (
	"log"
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
)

func RunRouters() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", homeHandleFunc).Methods("GET")
	r.HandleFunc("/products", productsHandleFunc).Methods("GET")
	return r
}

func homeHandleFunc(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles("views/html/index.html")
	if err != nil {
		log.Fatal(err)
	}
	w.WriteHeader(http.StatusOK)
	tpl.Execute(w, nil)
}

func productsHandleFunc(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles("views/html/products.html")
	if err != nil {
		log.Fatal(err)
	}
	w.WriteHeader(http.StatusOK)
	tpl.Execute(w, nil)
}
