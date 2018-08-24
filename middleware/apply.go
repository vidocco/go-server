package middleware

import (
	"github.com/gorilla/mux"
	"net/http"
	"log"
)

func logger(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s requested %s", r.RemoteAddr, r.URL)
		h.ServeHTTP(w, r)
	})
}

func ApplyMiddleware(router *mux.Router) http.Handler {
	handler := logger(router)
	return handler
}
