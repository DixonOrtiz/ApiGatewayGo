package database

import (
	"context"
	"fmt"

	"cloud.google.com/go/firestore"
	"github.com/DixonOrtiz/ApiGateway/api/functions"
)

var projectID = functions.GetEnv("PROJECT_ID")

//CreateFirestoreClient function
//Function that creates a new Firestore's Client
func CreateFirestoreClient() (*firestore.Client, context.Context, error) {
	ctx := context.Background()

	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		fmt.Println("failed to create a Firestore client")
		return nil, ctx, err
	}

	fmt.Printf("%T", ctx)

	return client, ctx, nil
}
