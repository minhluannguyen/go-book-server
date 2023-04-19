package routes

import (
	"github.com/gorilla/mux"
	"github.com/minhluannguyen/go-book-server/pkg/controllers"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/book", controllers.GetBook).Methods("GET")
	//router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
}
