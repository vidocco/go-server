package main

import (
	"net/http"
	"go-server/router"
	"go-server/middleware"
)

func main() {
	r := router.NewRouter()
	handler := middleware.ApplyMiddleware(r)
	// http.NewServeMux crea una especie de router https://cryptic.io/go-http/
	// para hacer "middlewares" creas funciones que reciben el anterior handler y devuelven un nuevo handler (como
	// en el ejemplo de arriba)
	err := http.ListenAndServe(":3000", handler)
	// el router se pasa al listenAndServe
	if err != nil {
		panic(err)
	}
}