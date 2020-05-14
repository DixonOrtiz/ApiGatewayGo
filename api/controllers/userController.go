package controllers

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/DixonOrtiz/ApiGateway/api/functions"
)

var userURL = functions.GetEnv("USER_URL")

//GetCurrentUser controller
//This controller gets the user by its GoogleID
func GetCurrentUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("[Gateway API][GET][USER][/user/currentUser}]")

	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalln(err)
	}

	endpoint := fmt.Sprintf("%s/user", userURL)

	response, err := http.Post(endpoint, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		log.Fatalf("failed getting user info: %s", err.Error())
	}
	defer response.Body.Close()

	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalf("failed getting user info: %s", err.Error())
	}

	fmt.Println("[Gateway API][GET][USER][/user/currentUser}][PASSED]")
	functions.JSON(w, http.StatusOK, string(responseBody))
}

//GetAllDevices controller
//This controller get all the devices in the database
func GetAllDevices(w http.ResponseWriter, r *http.Request) {
	fmt.Println("[Gateway API][GET][USER][/user/allDevices}]")

	endpoint := fmt.Sprintf("%s/allDevices", userURL)

	response, err := http.Get(endpoint)
	if err != nil {
		log.Fatalf("failed getting user info: %s", err.Error())
	}
	defer response.Body.Close()

	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalf("failed getting user info: %s", err.Error())

	}

	fmt.Println("[Gateway API][GET][USER][/user/allDevices}][PASSED]")
	functions.JSON(w, http.StatusOK, string(responseBody))
}
