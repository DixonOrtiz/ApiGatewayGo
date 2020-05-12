package api

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

//Run function
//Function that raises the server and define the endpoints
func Run() {
	router := mux.NewRouter()

	//Auth (Google OAuth2 and JWT) routes
	// router.HandleFunc("/login", controllers.HandleGoogleLogin)
	// router.HandleFunc("/callback", controllers.HandleGoogleCallback)

	//User routes

	e := godotenv.Load()
	if e != nil {
		fmt.Println(e)
		return
	}

	portEnv := os.Getenv("PORT")
	port := fmt.Sprintf(":%s", portEnv)

	fmt.Printf("Running in port %s\n", portEnv)
	http.ListenAndServe(port, router)

}
