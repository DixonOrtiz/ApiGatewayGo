package auth

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/DixonOrtiz/ApiGateway/api/functions"
	jwt "github.com/dgrijalva/jwt-go"
)

var jwtKey = functions.GetEnv("JWT_KEY")

//UserData struct to receive the google auth responsw
type UserData struct {
	TokenJWT string `json:"token"`
	GoogleID string `json:"id"`
	Name     string `json:"given_name"`
	Lastname string `json:"family_name"`
	Email    string `json:"email"`
	Photo    string `json:"picture"`
}

//CreateToken function
//This function create a new JWT
func CreateToken(user UserData) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["google_id"] = user.GoogleID
	claims["name"] = user.Name
	claims["lastname"] = user.Lastname
	claims["email"] = user.Email
	claims["photo"] = user.Photo
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(jwtKey))
}

//ExtractToken function
//This function extract de JWT from a request
func ExtractToken(r *http.Request) string {
	keys := r.URL.Query()
	token := keys.Get("token")

	if token != "" {
		return token
	}

	bearerToken := r.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}

//ExtractTokenGoogleID function
//This function extract de JWT from a request and return the googleID
func ExtractTokenGoogleID(r *http.Request) string {
	tokenString := ExtractToken(r)
	claims := jwt.MapClaims{}

	_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtKey), nil
	})
	if err != nil {
		fmt.Println(err)
	}

	googleID := claims["google_id"]

	return fmt.Sprint(googleID)
}

//TokenValidRequest function
//This function is responsible for validating the token
func TokenValidRequest(r *http.Request) error {
	tokenString := ExtractToken(r)

	_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(jwtKey), nil
	})
	if err != nil {
		return err
	}

	return nil
}

//TokenValid function
//This function is responsible for validating the token
func TokenValid(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(jwtKey), nil
	})
	if err != nil {
		return err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		Pretty(claims)
	}
	return nil
}

//Pretty function
//Display the claims licely in the terminal
func Pretty(data interface{}) {
	b, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println(string(b))
}
