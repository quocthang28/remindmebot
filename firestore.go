package main

import (
	"context"
	"log"
	"os"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

func NewFirebaseClient() (*firestore.Client, error) {
	log.Println("Connecting to Firebase...")

	ctx := context.Background()
	opt := option.WithCredentialsJSON(decrypt(os.Getenv("FIREBASE_CRED_K"), os.Getenv("FIREBASE_CRED")))

	firebaseApp, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	client, err := firebaseApp.Firestore(ctx)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return client, nil
}
