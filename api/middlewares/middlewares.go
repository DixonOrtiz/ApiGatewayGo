package middlewares

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/DixonOrtiz/ApiGateway/api/auth"
	"github.com/DixonOrtiz/ApiGateway/api/functions"
)

//UserAuthentication middleware
//This function verifies that the entered token is valid and contains a google user, if it exists it authorizes it
//and otherwise it registers it as a user without any associated device in firestore DB and authorizes it
func UserAuthentication(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("[Gateway API][Middleware][UserAuthentication]")

		//-> RECIBE EL JWT, Y LO VALIDA, SI ES VALIDO PASA, SI NO DESAUTORIZA
		err := auth.TokenValidRequest(r)
		if err != nil {
			fmt.Println("[Gateway API][Middleware][UserAuthentication][Unauthorized]")
			functions.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
			return
		}

		fmt.Println("googleID:", auth.ExtractTokenGoogleID(r))

		//-> CONSULTA EL USUARIO DE LA BASE DE DATOS (FIRESTORE DB)
		//-> SI NO EXISTE, LO CREA (FIRESTORE DB)

		fmt.Println("[Gateway API][Middleware][UserAuthentication][Authorized]")
		next(w, r)
	}
}

//ProtectedAuthentication middleware
//Verify that the consulted devices belong to the user who is consulting
func ProtectedAuthentication(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("[Gateway API][Middleware][ProtectedAuthentication]")

		//-> RECIBE EL JWT, Y LO VALIDA, SI ES VALIDO PASA, SI NO DESAUTORIZA

		//-> CONSULTA SI EL DISPOSITIVO CONSULTADO PERTENECE AL USUARIO TEL TOKEN (FIRESTORE DB)
		//-> SI NO LO ES DESAUTORIZA

		fmt.Println("[Gateway API][Middleware][ProtectedAuthentication][Authorized]")
		next(w, r)
	}
}

//AdminAuthentication middleware
//Verify that the consulted devices belong to the user who is consulting
func AdminAuthentication(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("[Gateway API][Middleware][AdminAuthentication]")

		//-> RECIBE EL JWT, Y LO VALIDA, SI ES VALIDO PASA, SI NO DESAUTORIZA

		//-> CONSULTA SI EL USUARIO ES UN ADMIN EN LA BASE DE DATOS (FIRESTORE DB)
		//-> SI NO LO ES DESAUTORIZA

		fmt.Println("[Gateway API][Middleware][AdminAuthentication][Authorized]")
		next(w, r)
	}
}
