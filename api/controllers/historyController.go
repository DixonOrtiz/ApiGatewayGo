package controllers

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/DixonOrtiz/ApiGateway/api/functions"
	"github.com/gorilla/mux"
)

var historyURL = functions.GetEnv("HISTORY_URL")

//GetDayGraph controller
//This controllers gets device's day history
func GetDayGraph(w http.ResponseWriter, r *http.Request) {
	deviceID := mux.Vars(r)["deviceID"]
	fmt.Printf("[Gateway API][Get][History][/history/day/%s}]\n", deviceID)

	endpoint := fmt.Sprintf("%s/day/%s", historyURL, deviceID)

	response, err := http.Get(endpoint)
	if err != nil {
		fmt.Printf("failed getting device's day history by its deviceID: %s\n", err.Error())
		functions.ERROR(w, http.StatusBadRequest, errors.New("bad request"))
		return
	}
	defer response.Body.Close()

	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("failed getting device's day history by its deviceID: %s\n", err.Error())
		functions.ERROR(w, http.StatusBadRequest, errors.New("bad request"))
		return
	}

	fmt.Printf("[Gateway API][Get][History][/history/day/%s}]\n][Response]\n", deviceID)
	functions.PrettyJSONTerminal([]byte(responseBody))
	functions.ResponseJSON(w, http.StatusOK, string(responseBody))
}

//GetWeekGraph controller
//This controllers gets device's week history
func GetWeekGraph(w http.ResponseWriter, r *http.Request) {
	deviceID := mux.Vars(r)["deviceID"]
	fmt.Printf("[Gateway API][Get][History][/history/week/%s}]\n", deviceID)

	endpoint := fmt.Sprintf("%s/week/%s", historyURL, deviceID)
	response, err := http.Get(endpoint)
	if err != nil {
		fmt.Printf("failed getting device's week history by its deviceID: %s\n", err.Error())
		functions.ERROR(w, http.StatusBadRequest, errors.New("bad request"))
		return
	}
	defer response.Body.Close()

	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("failed getting device's week history by its deviceID: %s\n", err.Error())
		functions.ERROR(w, http.StatusBadRequest, errors.New("bad request"))
		return
	}

	fmt.Printf("[Gateway API][Get][History][/history/week/%s}]\n][Response]\n", deviceID)
	functions.PrettyJSONTerminal([]byte(responseBody))
	functions.ResponseJSON(w, http.StatusOK, string(responseBody))
}

//GetMonthGraph controller
//This controllers gets device's month history
func GetMonthGraph(w http.ResponseWriter, r *http.Request) {
	deviceID := mux.Vars(r)["deviceID"]
	fmt.Printf("[Gateway API][Get][History][/history/month/%s}]\n", deviceID)

	endpoint := fmt.Sprintf("%s/month/%s", historyURL, deviceID)
	response, err := http.Get(endpoint)
	if err != nil {
		fmt.Printf("failed getting device's month history by its deviceID: %s\n", err.Error())
		functions.ERROR(w, http.StatusBadRequest, errors.New("bad request"))
		return

	}
	defer response.Body.Close()

	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("failed getting device's month history by its deviceID: %s\n", err.Error())
		functions.ERROR(w, http.StatusBadRequest, errors.New("bad request"))
		return
	}

	fmt.Printf("[Gateway API][Get][History][/history/month/%s}]\n][Response]\n", deviceID)
	functions.PrettyJSONTerminal([]byte(responseBody))
	functions.ResponseJSON(w, http.StatusOK, string(responseBody))
}
