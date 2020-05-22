package controllers

import (
	"errors"
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
		fmt.Printf("failed getting device by its deviceID: %s\n", err.Error())
		functions.ERROR(w, http.StatusBadRequest, errors.New("bad request"))
		return
	}

	defer response.Body.Close()

	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("failed getting device by its deviceID: %s\n", err.Error())
		functions.ERROR(w, http.StatusBadRequest, errors.New("bad request"))
		return
	}

	fmt.Printf("[Gateway API][Get][Device-Control][/deviceControl/device/%s/state}]\n][Response]\n", deviceID)
	functions.PrettyJSONTerminal([]byte(responseBody))
	functions.ResponseJSON(w, http.StatusOK, string(responseBody))
}

//GetDeviceLastConfig controller
//This controllers gets the last device's config
func GetDeviceLastConfig(w http.ResponseWriter, r *http.Request) {
	deviceID := mux.Vars(r)["deviceID"]
	fmt.Printf("[Gateway API][Get][Device-Control][/deviceControl/device/%s/config}]\n", deviceID)

	endpoint := fmt.Sprintf("%s/device/%s/config", deviceControlURL, deviceID)

	response, err := http.Get(endpoint)
	if err != nil {
		fmt.Printf("failed getting device last config: %s\n", err.Error())
		functions.ERROR(w, http.StatusBadRequest, errors.New("bad request"))
		return

	}
	defer response.Body.Close()

	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("failed getting device las last config: %s", err.Error())
		functions.ERROR(w, http.StatusBadRequest, errors.New("bad request"))
		return
	}

	fmt.Printf("[Gateway API][Get][Device-Control][/deviceControl/device/%s/config}]\n][Response]\n", deviceID)
	functions.PrettyJSONTerminal([]byte(responseBody))
	functions.ResponseJSON(w, http.StatusOK, string(responseBody))
}

//GetDeviceHistoryState controller
//This controllers gets the device's history state
func GetDeviceHistoryState(w http.ResponseWriter, r *http.Request) {
	deviceID := mux.Vars(r)["deviceID"]
	fmt.Printf("[Gateway API][Get][Device-Control][/deviceControl/device/%s/stateHistory}]\n", deviceID)

	endpoint := fmt.Sprintf("%s/device/%s/state-history", deviceControlURL, deviceID)

	response, err := http.Get(endpoint)
	if err != nil {
		fmt.Printf("failed getting device's history state by its deviceID: %sn", err.Error())
		functions.ERROR(w, http.StatusBadRequest, errors.New("bad request"))
		return
	}
	defer response.Body.Close()

	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("failed getting device's history state by its deviceID: %s\n", err.Error())
		functions.ERROR(w, http.StatusBadRequest, errors.New("bad request"))
		return
	}

	fmt.Printf("[Gateway API][Get][Device-Control][/deviceControl/device/%s/stateHistory}]\n][Response]\n", deviceID)
	functions.PrettyJSONTerminal([]byte(responseBody))
	functions.ResponseJSON(w, http.StatusOK, string(responseBody))
}

//GetDeviceHistoryConfig controller
//This controllers gets the device's config state
func GetDeviceHistoryConfig(w http.ResponseWriter, r *http.Request) {
	deviceID := mux.Vars(r)["deviceID"]
	fmt.Printf("[Gateway API][Get][Device-Control][/deviceControl/device/%s/configHistory}]\n", deviceID)

	endpoint := fmt.Sprintf("%s/device/%s/config-history", deviceControlURL, deviceID)

	response, err := http.Get(endpoint)
	if err != nil {
		fmt.Printf("failed getting device's history config by its deviceID: %s\n", err.Error())
		functions.ERROR(w, http.StatusBadRequest, errors.New("bad request"))
		return
	}
	defer response.Body.Close()

	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("failed getting device's history config by its deviceID: %s\n", err.Error())
		functions.ERROR(w, http.StatusBadRequest, errors.New("bad request"))
		return
	}

	fmt.Printf("[Gateway API][Get][Device-Control][/deviceControl/device/%s/configHistory}]\n][Response]\n", deviceID)
	functions.PrettyJSONTerminal([]byte(responseBody))
	functions.ResponseJSON(w, http.StatusOK, string(responseBody))
}

//GetRegistries controller
//This controllers list the registries
func GetRegistries(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("[Gateway API][Get][Device-Control][/deviceControl/registry]\n")

	endpoint := fmt.Sprintf("%s/registry", deviceControlURL)

	response, err := http.Get(endpoint)
	if err != nil {
		log.Fatalf("failed getting device by its deviceID: %s", err.Error())
	}
	defer response.Body.Close()

	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("failed getting device by its deviceID: %s\n", err.Error())
		functions.ERROR(w, http.StatusBadRequest, errors.New("bad request"))
		return
	}

	fmt.Printf("[Gateway API][Get][Device-Control][/deviceControl/registry]\n][Response]\n")
	functions.PrettyJSONTerminal([]byte(responseBody))
	functions.ResponseJSON(w, http.StatusOK, string(responseBody))
}
