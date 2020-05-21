package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/DixonOrtiz/ApiGateway/api/auth"
	"github.com/DixonOrtiz/ApiGateway/api/functions"
)

//UserData struct to receive the google auth response
type UserData struct {
	TokenJWT string `json:"token"`
	GoogleID string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Photo    string `json:"picture"`
}

//HandleGoogleLogin controler
func HandleGoogleLogin(w http.ResponseWriter, r *http.Request) {
	fmt.Println("[Gateway API][GET][AUTH][/login]")

	url := auth.GoogleOauthConfig.AuthCodeURL(auth.OauthStateString)

	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
	fmt.Println("[Gateway API][GET][AUTH][/login][PASSED]")
}

//HandleGoogleCallback controller
func HandleGoogleCallback(w http.ResponseWriter, r *http.Request) {
	fmt.Println("[Gateway API][GET][AUTH][/callback]")
	content, err := auth.GetUserInfo(r.FormValue("state"), r.FormValue("code"))
	if err != nil {
		fmt.Println(err.Error())
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	user := &UserData{}

	err = json.Unmarshal(content, user) //paso de data de un []bytes a un struct de go
	if err != nil {
		fmt.Println(err)
	}

	user.TokenJWT, err = auth.CreateToken(user.GoogleID)
	if err != nil {
		fmt.Println(user.TokenJWT)
	}

	fmt.Println()

	fmt.Println("[Gateway API][GET][AUTH][/callback][PASSED]")
	functions.ResponseLoginJSON(w, http.StatusOK, user)
}
