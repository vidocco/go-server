package router

import (
	"github.com/naoina/denco"
	"net/http"
	"go-server/controllers"
)

func NewRouter() http.Handler {
	mux := denco.NewMux()

	router, err := mux.Build([]denco.Handler{
		mux.GET("/users", controllers.GetAllUsers),
		mux.GET("/users/:userId", controllers.GetUser),
		mux.POST("/users", controllers.PostUser),
	})

	if err != nil {
		panic(err)
	}
	return router
}