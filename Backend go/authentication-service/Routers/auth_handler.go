package routers

import (
	"net/http"
	"task-manager-microservices/authentication-service/services"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func Router(){

	services.SqlConnection()
	router := mux.NewRouter()

	corsHandler := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
	)

	router.HandleFunc("/registartion", services.PostRegistartionContorller).Methods("POST")
	router.HandleFunc("/getAll", services.GetAllRegisteredUsersController).Methods("GET")
	router.HandleFunc("/loggedIn", services.CheckLogInDetails).Methods("POST")

	http.ListenAndServe(":4000", corsHandler(router))
	
}


