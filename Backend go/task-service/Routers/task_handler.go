package routers

import (
	"net/http"
	"task-manager-microservices/task-service/services"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func Route()  {

	services.SqlConnection()
	router := mux.NewRouter()

	corsHandler := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
	)

	router.HandleFunc("/postTask", services.AddTask). Methods("POST")
	router.HandleFunc("/getTaskForUser/{username}", services.GetTaskForUser).Methods("GET")
	router.HandleFunc("/getTaskForCreatedUser/{username}", services.GetTaskCreatedUser).Methods("GET")
	router.HandleFunc("/updateStatus/{username}/{title}/{status}", services.ChangeTaskStatus).Methods("PUT")
	router.HandleFunc("/updateStatusCreatedUser/{username}/{title}/{status}", services.ChangeTaskStatusCreated).Methods("PUT")

	http.ListenAndServe(":4001", corsHandler(router))

}
