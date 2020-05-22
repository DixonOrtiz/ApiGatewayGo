package controllers

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/DixonOrtiz/ApiGateway/api/functions"
	"github.com/gorilla/mux"
)

var userURL = functions.GetEnv("USER_URL")

//GetCurrentUser controller
//This controller gets the user by its GoogleID
func GetCurrentUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("[Gateway API][GET][USER][/user/currentUser}]")

	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("failed getting the current user: %s\n", err.Error())
		functions.ERROR(w, http.StatusBadRequest, errors.New("bad request"))
		return
	}

	endpoint := fmt.Sprintf("%s/user", userURL)

	response, err := http.Post(endpoint, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		fmt.Printf("failed getting the current user: %s\n", err.Error())
		functions.ERROR(w, http.StatusBadRequest, errors.New("bad request"))
		return
	}
	defer response.Body.Close()

	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("failed getting the current user: %s\n", err.Error())
		functions.ERROR(w, http.StatusBadRequest, errors.New("bad request"))
		return
	}

	fmt.Println("[Gateway API][GET][USER][/user/currentUser}][RESPONSE]")
	functions.PrettyJSONTerminal(responseBody)
	functions.ResponseJSON(w, http.StatusOK, string(responseBody))
}

//GetDevices controller
//This controller gets the user's devices by its GoogleID
func GetDevices(w http.ResponseWriter, r *http.Request) {
	fmt.Println("[Gateway API][GET][USER][/user/devices}]")

	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("failed getting the users' devices: %s\n", err.Error())
		functions.ERROR(w, http.StatusBadRequest, errors.New("bad request"))
		return
	}

	endpoint := fmt.Sprintf("%s/devices", userURL)

	response, err := http.Post(endpoint, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		fmt.Printf("failed getting the users' devices: %s\n", err.Error())
		functions.ERROR(w, http.StatusBadRequest, errors.New("bad request"))
		return
	}
	defer response.Body.Close()

	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("failed getting the users' devices: %s\n", err.Error())
		functions.ERROR(w, http.StatusBadRequest, errors.New("bad request"))
		return
	}

	fmt.Println("[Gateway API][GET][USER][/user/devices}][RESPONSE]")
	functions.PrettyJSONTerminal(responseBody)
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
		fmt.Printf("failed posting a new device: %s\n", err.Error())
		functions.ERROR(w, http.StatusBadRequest, errors.New("bad request"))
		return
	}

	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("failed posting a new device: %s\n", err.Error())
		functions.ERROR(w, http.StatusBadRequest, errors.New("bad request"))
		return
	}
	defer response.Body.Close()

	fmt.Println("[Gateway API][POST][USER][/user/saveDevice}][RESPONSE]")
	functions.PrettyJSONTerminal(responseBody)
	functions.ResponseJSON(w, http.StatusOK, string(responseBody))
}

//ChangeDevice controller
//This controller change an existing device by its deviceId
func ChangeDevice(w http.ResponseWriter, r *http.Request) {
	fmt.Println("[Gateway API][POST][USER][/user/changeDevice}]")

	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalln(err)
	}

	endpoint := fmt.Sprintf("%s/changeDevice", userURL)

	response, err := http.Post(endpoint, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		fmt.Printf("failed changing a device: %s\n", err.Error())
		functions.ERROR(w, http.StatusBadRequest, errors.New("bad request"))
		return
	}

	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("failed changing a device: %s\n", err.Error())
		functions.ERROR(w, http.StatusBadRequest, errors.New("bad request"))
		return
	}
	defer response.Body.Close()

	fmt.Println("[Gateway API][POST][USER][/user/changeDevice}][RESPONSE]")
	functions.PrettyJSONTerminal(responseBody)
	functions.ResponseJSON(w, http.StatusOK, string(responseBody))
}

//GetAllDevices controller
//This controller get all the devices in the database
func GetAllDevices(w http.ResponseWriter, r *http.Request) {
	fmt.Println("[Gateway API][GET][USER][/user/allDevices}]")

	endpoint := fmt.Sprintf("%s/allDevices", userURL)

	response, err := http.Get(endpoint)
	if err != nil {
		fmt.Printf("failed getting all devices: %s\n", err.Error())
		functions.ERROR(w, http.StatusBadRequest, errors.New("bad request"))
		return
	}
	defer response.Body.Close()

	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("failed getting all devices: %s\n", err.Error())
		functions.ERROR(w, http.StatusBadRequest, errors.New("bad request"))
		return
	}

	fmt.Println("[Gateway API][Get][User][/user/allDevices}][Response]")
	functions.PrettyJSONTerminal(responseBody)
	functions.ResponseJSON(w, http.StatusOK, string(responseBody))
}

//GetUserByDevice controller
//Thsi controller gets the data info with a deviceId
func GetUserByDevice(w http.ResponseWriter, r *http.Request) {
	deviceID := mux.Vars(r)["deviceID"]
	fmt.Printf("[Gateway API][Get][User][/user/device/%s/user}]\n", deviceID)

	endpoint := fmt.Sprintf("%s/device/%s/user", userURL, deviceID)

	fmt.Println(endpoint)

	response, err := http.Get(endpoint)
	if err != nil {
		fmt.Printf("failed getting user by its device: %s\n", err.Error())
		functions.ERROR(w, http.StatusBadRequest, errors.New("bad request"))
		return
	}
	defer response.Body.Close()

	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("failed getting user by its device: %s\n", err.Error())
		functions.ERROR(w, http.StatusBadRequest, errors.New("bad request"))
		return
	}

	fmt.Printf("[Gateway API][Get][User][/user/device/%s/user}][Response]\n", deviceID)
	functions.PrettyJSONTerminal([]byte(responseBody))
	functions.ResponseJSON(w, http.StatusOK, string(responseBody))
}
