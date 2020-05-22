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
	router.HandleFunc("/user/currentUser", middlewares.UserAuthentication(controllers.GetCurrentUser)).Methods("GET")
	router.HandleFunc("/user/devices", controllers.GetDevices).Methods("GET")
	router.HandleFunc("/user/saveDevice", middlewares.UserAuthentication(controllers.SaveDevice)).Methods("POST")
	router.HandleFunc("/user/changeDevice", controllers.ChangeDevice).Methods("POST")
	router.HandleFunc("/user/allDevices", middlewares.AdminAuthentication(controllers.GetAllDevices)).Methods("GET")
	router.HandleFunc("/user/device/{deviceID}/user", controllers.GetUserByDevice).Methods("GET")

	//Device Control routes
	router.HandleFunc("/deviceControl/device/{deviceID}/state", controllers.GetDeviceLastState).Methods("GET")
	router.HandleFunc("/deviceControl/device/{deviceID}/config", controllers.GetDeviceLastConfig).Methods("GET")
	// router.HandleFunc("/deviceControl/device/{deviceID}", controllers.updateDeviceConfig).Methods("PUT")
	router.HandleFunc("/deviceControl/device/{deviceID}/stateHistory", controllers.GetDeviceHistoryState).Methods("GET")
	router.HandleFunc("/deviceControl/device/{deviceID}/configHistory", controllers.GetDeviceHistoryConfig).Methods("GET")
	router.HandleFunc("/deviceControl/registry", controllers.GetRegistries).Methods("GET")

	//Device Control routes
	router.HandleFunc("/history/day/{deviceID}", middlewares.ProtectedAuthentication(controllers.GetDayGraph)).Methods("GET")
	router.HandleFunc("/history/week/{deviceID}", controllers.GetWeekGraph).Methods("GET")
	router.HandleFunc("/history/month/{deviceID}", controllers.GetMonthGraph).Methods("GET")

	fmt.Printf("Running in port %s\n", portEnv)
	http.ListenAndServe(port, router)

}
