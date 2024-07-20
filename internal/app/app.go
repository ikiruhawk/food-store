package app

import (
	"net/http"

	"github.com/ikiruhawk/food-store/internal/routers"
)

func Run() {
	r := routers.RunRouters()
	http.ListenAndServe(":8080", r)
}
