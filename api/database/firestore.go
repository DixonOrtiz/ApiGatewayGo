package database

import (
	"fmt"
	"log"
)

//User type collection in Firstore
type User struct {
	GoogleID string `json:"id"`
	Name     string `json:"given_name"`
	Lastname string `json:"family_name"`
	Email    string `json:"email"`
	Photo    string `json:"picture"`
}

//CreateUser function
//Function that creates a new user in Firestore Database
func CreateUser(user *User) (*User, error) {
	client, ctx, err := CreateFirestoreClient()
	if err != nil {
		fmt.Println("failed to create a Firestore client")
	}
	defer client.Close()

	_, _, err = client.Collection("users").Add(ctx, map[string]interface{}{
		"googleID": user.GoogleID,
		"name":     user.Name,
		"lastName": user.Lastname,
		"email":    user.Email,
		"photo":    user.Photo,
	})
	if err != nil {
		log.Fatalf("failed adding a new user: %v", err)
		return nil, err
	}
	return user, nil

}
