package api

import (
	"fmt"
	"net/http"

	"github.com/DixonOrtiz/ApiGateway/api/controllers"
	"github.com/DixonOrtiz/ApiGateway/api/functions"
	"github.com/DixonOrtiz/ApiGateway/api/middlewares"
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
	router.HandleFunc("/user/currentUser", controllers.GetCurrentUser).Methods("GET")
	router.HandleFunc("/user/devices", controllers.GetDevices).Methods("GET")
	router.HandleFunc("/user/saveDevice", middlewares.UserAuthentication(controllers.SaveDevice)).Methods("POST")
	router.HandleFunc("/user/changeDevice", middlewares.UserAuthentication(controllers.ChangeDevice)).Methods("POST")
	router.HandleFunc("/user/allDevices", middlewares.AdminAuthentication(controllers.GetAllDevices)).Methods("GET")
	router.HandleFunc("/user/device/{deviceID}/user", controllers.GetUserByDevice).Methods("GET")

	//Device Control routes
	router.HandleFunc("/deviceControl/device/{deviceID}/state", middlewares.ProtectedAuthentication(controllers.GetDeviceLastState)).Methods("GET")
	router.HandleFunc("/deviceControl/device/{deviceID}/config", middlewares.ProtectedAuthentication(controllers.GetDeviceLastConfig)).Methods("GET")
	router.HandleFunc("/deviceControl/device/{deviceID}/stateHistory", middlewares.ProtectedAuthentication(controllers.GetDeviceHistoryState)).Methods("GET")
	router.HandleFunc("/deviceControl/device/{deviceID}/configHistory", middlewares.ProtectedAuthentication(controllers.GetDeviceHistoryConfig)).Methods("GET")
	router.HandleFunc("/deviceControl/registry", middlewares.AdminAuthentication(controllers.GetRegistries)).Methods("GET")

	//Device Control routes
	router.HandleFunc("/history/day/{deviceID}", middlewares.ProtectedAuthentication(controllers.GetDayGraph)).Methods("GET")
	router.HandleFunc("/history/week/{deviceID}", middlewares.ProtectedAuthentication(controllers.GetWeekGraph)).Methods("GET")
	router.HandleFunc("/history/month/{deviceID}", middlewares.ProtectedAuthentication(controllers.GetMonthGraph)).Methods("GET")

	fmt.Printf("Running in port %s\n", portEnv)
	http.ListenAndServe(port, router)

}
