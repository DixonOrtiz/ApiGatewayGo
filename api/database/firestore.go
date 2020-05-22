package database

import (
	"fmt"
	"log"

	"github.com/mitchellh/mapstructure"
	"google.golang.org/api/iterator"
)

//User type collection in Firstore
type User struct {
	GoogleID string
	Name     string
	Lastname string
	Email    string
	Photo    string
}

//Device type collection in Firstore
type Device struct {
	DeviceID string
	User     string //GoogleID in device document
}

//GetUser function
//Function that gets an user in Firestore Database
func GetUser(googleID string) (*User, bool, error) {
	client, ctx, err := CreateFirestoreClient()
	if err != nil {
		fmt.Println("failed to create a Firestore client")
	}
	defer client.Close()

	query := client.Collection("users").Where("googleID", "==", googleID).Documents(ctx)
	for {
		doc, err := query.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return &User{}, false, err
		}
		m := map[string]interface{}(doc.Data())

		user := &User{}
		mapstructure.Decode(m, &user)

		return user, true, nil
	}

	return &User{}, false, err
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

//VerifyDeviceUser function
//Function that verifies if the deviceID and the googleID matches
func VerifyDeviceUser(deviceID, googleID string) (bool, error) {
	client, ctx, err := CreateFirestoreClient()
	if err != nil {
		fmt.Println("failed to create a Firestore client")
	}
	defer client.Close()

	query := client.Collection("devices").Where("deviceID", "==", deviceID).Documents(ctx)
	for {
		doc, err := query.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return false, err
		}

		m := map[string]interface{}(doc.Data())

		device := &Device{}
		mapstructure.Decode(m, &device)

		if device.User == googleID {
			return true, nil
		}
	}

	return false, nil
}
