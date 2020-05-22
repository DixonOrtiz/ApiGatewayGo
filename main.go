package main

import (
	// "github.com/DixonOrtiz/ApiGateway/api"s
	"github.com/DixonOrtiz/ApiGateway/api/database"
)

func main() {
	user := database.User{
		GoogleID: "test-googleid",
		Name:     "test-name",
		Lastname: "test-lastname",
		Email:    "test-email",
		Photo:    "test-photo",
	}

	database.CreateUser(&user)
}
