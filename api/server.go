package api

import (
	"fmt"
	"net/http"

	"github.com/DixonOrtiz/ApiGateway/api/controllers"
	"github.com/DixonOrtiz/ApiGateway/api/functions"
	"github.com/gorilla/mux"
)

//Run function
//Function that raises the server and define the endpoints
func Run() {
	router := mux.NewRouter()

	//Auth (Google OAuth2 and JWT) routes
	router.HandleFunc("/login", controllers.HandleGoogleLogin)
	router.HandleFunc("/callback", controllers.HandleGoogleCallback)

	//User routes

	portEnv := functions.GetEnv("PORT")
	port := fmt.Sprintf(":%s", portEnv)

	fmt.Printf("Running in port %s\n", portEnv)
	http.ListenAndServe(port, router)

}
