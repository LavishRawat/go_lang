package routes

import(
	"github.com/gorilla/mux"
	controller "Server/src/controllers"
	// "log"
	// "net/http"
)

var RegisterSignUpDetailsRoutes = func(router *mux.Router){
	router.HandleFunc("/signup",controller.StoreSignupDetails).Methods("POST")
}

