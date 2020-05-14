package controllers

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/DixonOrtiz/ApiGateway/api/functions"
)

//GetAllDevices controller
//This controller make an http get request to User Api
func GetAllDevices(w http.ResponseWriter, r *http.Request) {
	fmt.Println("[Gateway API][GET][USER][/device/allDevices}]")

	userURL := functions.GetEnv("USER_URL")
	endpoint := fmt.Sprintf("%s/allDevices", userURL)

	response, err := http.Get(endpoint)
	if err != nil {
		fmt.Printf("failed getting user info: %s", err.Error())
	}

	defer response.Body.Close()

	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("failed getting user info: %s", err.Error())

	}

	fmt.Println("[Gateway API][GET][USER][/device/allDevices}][PASSED]")

	functions.JSON(w, http.StatusOK, string(responseBody))
}
