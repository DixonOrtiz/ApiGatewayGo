package middlewares

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/DixonOrtiz/ApiGateway/api/auth"
	"github.com/DixonOrtiz/ApiGateway/api/database"
	"github.com/DixonOrtiz/ApiGateway/api/functions"
	"github.com/gorilla/mux"
)

//UserAuthentication middleware
//This function verifies that the entered token is valid and contains a google user, if it exists it authorizes it
//and otherwise it registers it as a user without any associated device in firestore DB and authorizes it
func UserAuthentication(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("[Gateway API][Middleware][UserAuthentication]")

		err := auth.TokenValidRequest(r)
		if err != nil {
			fmt.Println("[Gateway API][Middleware][UserAuthentication][Unauthorized]")
			functions.ERROR(w, http.StatusUnauthorized, errors.New("TokenValidRequest error"))
			return
		}

		googleID := auth.ExtractTokenGoogleID(r)

		userFirestore, boolUserExists, err := database.GetUser(googleID)
		if err != nil {
			fmt.Println("[Gateway API][Middleware][UserAuthentication][Unauthorized]")
			functions.ERROR(w, http.StatusUnauthorized, errors.New("ExtractTokenGoogleID error"))
			return
		}

		if !boolUserExists {
			user := auth.ExtractUser(r)
			database.CreateUser(&user)

			fmt.Println("[Gateway API][Middleware][UserAuthentication][User Created]", user)
		}

		fmt.Println("[Gateway API][Middleware][UserAuthentication][Authorized][User]", userFirestore)
		next(w, r)
	}
}

//ProtectedAuthentication middleware
//Verify that the consulted devices belong to the user who is consulting
func ProtectedAuthentication(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("[Gateway API][Middleware][ProtectedAuthentication]")

		err := auth.TokenValidRequest(r)
		if err != nil {
			fmt.Println("[Gateway API][Middleware][ProtectedAuthentication][Unauthorized]")
			functions.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
			return
		}

		deviceID := mux.Vars(r)["deviceID"]
		googleID := auth.ExtractTokenGoogleID(r)

		deviceUserMatch, err := database.VerifyDeviceUser(deviceID, googleID)

		if err != nil {
			fmt.Println("[Gateway API][Middleware][ProtectedAuthentication][Unauthorized]")
			functions.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
			return
		}

		if deviceUserMatch {
			fmt.Println("[Gateway API][Middleware][ProtectedAuthentication][Authorized]")
			next(w, r)
			return
		}

		fmt.Println("[Gateway API][Middleware][ProtectedAuthentication][Unauthorized]")
		functions.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}
}

//AdminAuthentication middleware
//Verify that the consulted devices belong to the user who is consulting
func AdminAuthentication(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("[Gateway API][Middleware][AdminAuthentication]")

		err := auth.TokenValidRequest(r)
		if err != nil {
			fmt.Println("[Gateway API][Middleware][AdminAuthentication][Unauthorized]")
			functions.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
			return
		}

		googleID := auth.ExtractTokenGoogleID(r)

		isAdmin, err := database.VerifyAdmin(googleID)
		if err != nil {
			fmt.Println("[Gateway API][Middleware][AdminAuthentication][Unauthorized]")
			functions.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
			return
		}

		if isAdmin {
			fmt.Println("[Gateway API][Middleware][AdminAuthentication][Authorized]")
			next(w, r)
			return
		}

		fmt.Println("[Gateway API][Middleware][AdminAuthentication][Unauthorized]")
		functions.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}
}
