package controllers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/DixonOrtiz/ApiGateway/api/functions"
	"github.com/gorilla/mux"
)

var deviceControlURL = functions.GetEnv("DEVICE_CONTROL_URL")

//GetDeviceLastState controller
//This controllers gets the last device's state
func GetDeviceLastState(w http.ResponseWriter, r *http.Request) {
	deviceID := mux.Vars(r)["deviceID"]
	fmt.Printf("[Gateway API][Get][Device-Control][/deviceControl/device/%s/state}]\n", deviceID)

	endpoint := fmt.Sprintf("%s/device/%s/state", deviceControlURL, deviceID)

	response, err := http.Get(endpoint)
	if err != nil {
		log.Fatalf("failed getting device by its deviceID: %s", err.Error())
	}
	defer response.Body.Close()

	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalf("failed getting device by its deviceID: %s", err.Error())

	}

	fmt.Printf("[Gateway API][Get][Device-Control][/deviceControl/device/%s/state}]\n][Response]\n", deviceID)
	functions.PrettyJSONTerminal([]byte(responseBody))
	functions.ResponseJSON(w, http.StatusOK, string(responseBody))
}
