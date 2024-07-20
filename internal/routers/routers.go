package routers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func RunRouters() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", homeHandleFunc)
	r.HandleFunc("/products", productsHandleFunc)
	return r
}

func homeHandleFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func productsHandleFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "List of products")
}
