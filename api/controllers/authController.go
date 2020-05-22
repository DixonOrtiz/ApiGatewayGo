package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/DixonOrtiz/ApiGateway/api/auth"
	"github.com/DixonOrtiz/ApiGateway/api/functions"
)

//HandleGoogleLogin controler
func HandleGoogleLogin(w http.ResponseWriter, r *http.Request) {
	fmt.Println("[Gateway API][Get][Auth][/login]")

	url := auth.GoogleOauthConfig.AuthCodeURL(auth.OauthStateString)

	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
	fmt.Println("[Gateway API][Get][Auth][/login][Passed]")
}

//HandleGoogleCallback controller
func HandleGoogleCallback(w http.ResponseWriter, r *http.Request) {
	fmt.Println("[Gateway API][Get][Auth][/callback]")
	content, err := auth.GetUserInfo(r.FormValue("state"), r.FormValue("code"))
	if err != nil {
		fmt.Println(err.Error())
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	user := &auth.UserData{}

	err = json.Unmarshal(content, user)
	if err != nil {
		fmt.Println(err)
	}

	user.TokenJWT, err = auth.CreateToken(*user)
	if err != nil {
		fmt.Println(user.TokenJWT)
	}

	fmt.Println("[Gateway API][Get][Auth][/callback][Passed]")
	functions.ResponseLoginJSON(w, http.StatusOK, user)
}
