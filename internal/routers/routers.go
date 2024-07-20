package routers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func RunRouters() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", homeHandleFunc).Methods("GET")
	r.HandleFunc("/products", productsHandleFunc).Methods("GET")
	return r
}

func homeHandleFunc(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Home Page")
}

func productsHandleFunc(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "List of products")
}
