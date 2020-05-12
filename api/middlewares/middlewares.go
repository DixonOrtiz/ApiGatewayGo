package middlewares

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/DixonOrtiz/ApiGateway/api/auth"
	"github.com/DixonOrtiz/ApiGateway/api/functions"
)

//SetMiddlewareAuthentication middleware
//This function verify that the jwt introduced is valid
func SetMiddlewareAuthentication(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("[Gateway API][MIDDLEWARE][SetMiddlewareAuthentication]")

		err := auth.TokenValidRequest(r)

		if err != nil {
			functions.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
			return
		}
		fmt.Println("[Gateway API][MIDDLEWARE][SetMiddlewareAuthentication][PASSED]")
		next(w, r)
	}
}
