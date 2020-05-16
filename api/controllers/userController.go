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
	functions.ResponseJSON(w, http.StatusOK, string(responseBody))
}

//GetDevices controller
//This controller gets the user's devices by its GoogleID
func GetDevices(w http.ResponseWriter, r *http.Request) {
	fmt.Println("[Gateway API][GET][USER][/user/devices}]")

	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalln(err)
	}

	endpoint := fmt.Sprintf("%s/devices", userURL)

	response, err := http.Post(endpoint, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		log.Fatalf("failed getting user's devices: %s", err.Error())
	}
	defer response.Body.Close()

	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalf("failed getting user's devices: %s", err.Error())
	}

	fmt.Println("[Gateway API][GET][USER][/user/devices}][PASSED]")
	functions.ResponseJSON(w, http.StatusOK, string(responseBody))
}

//SaveDevice controller
//This controller post a new device to an existing user
func SaveDevice(w http.ResponseWriter, r *http.Request) {
	fmt.Println("[Gateway API][POST][USER][/user/saveDevice}]")

	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalln(err)
	}

	endpoint := fmt.Sprintf("%s/saveDevice", userURL)

	response, err := http.Post(endpoint, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		log.Fatalf("failed posting a new device: %s", err.Error())
	}

	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalf("failed getting user's devices: %s", err.Error())
	}

	defer response.Body.Close()

	fmt.Println("[Gateway API][POST][USER][/user/saveDevice}][PASSED]")
	functions.ResponseJSON(w, http.StatusOK, string(responseBody))
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
	functions.ResponseJSON(w, http.StatusOK, string(responseBody))
}
