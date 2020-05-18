package api

import (
	"fmt"
	"net/http"

	"github.com/DixonOrtiz/ApiGateway/api/controllers"
	"github.com/DixonOrtiz/ApiGateway/api/functions"
	"github.com/gorilla/mux"
)

//Variables to call the "http.ListenAndServe(port, router)"
//function in "Run()" function defined in this package
var (
	portEnv = functions.GetEnv("PORT")
	port    = fmt.Sprintf(":%s", portEnv)
)

//Run function
//Function that raises the server and define the endpoints
func Run() {
	router := mux.NewRouter()

	//Auth (Google OAuth2 and JWT) routes
	router.HandleFunc("/login", controllers.HandleGoogleLogin)
	router.HandleFunc("/callback", controllers.HandleGoogleCallback)

	//User routes
	router.HandleFunc("/user/currentUser", controllers.GetCurrentUser).Methods("GET") //add JWT middleware
	router.HandleFunc("/user/devices", controllers.GetDevices).Methods("GET")         //add JWT middleware
	router.HandleFunc("/user/saveDevice", controllers.SaveDevice).Methods("POST")     //add JWT middleware
	router.HandleFunc("/user/changeDevice", controllers.ChangeDevice).Methods("POST")
	router.HandleFunc("/user/allDevices", controllers.GetAllDevices).Methods("GET")               //add admin middleware
	router.HandleFunc("/user/device/{deviceID}/user", controllers.GetUserByDevice).Methods("GET") //add admin middleware

	fmt.Printf("Running in port %s\n", portEnv)
	http.ListenAndServe(port, router)

}
