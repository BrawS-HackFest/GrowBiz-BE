package config

import (
	"context"
	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"google.golang.org/api/option"
	"log"
	"os"
)

func FirebaseApp() *firebase.App {
	opt := option.WithCredentialsFile("config/key.json")
	config := &firebase.Config{ProjectID: os.Getenv("PROJECT_ID")}
	app, err := firebase.NewApp(context.Background(), config, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n\n", err)
	}
	return app
}

func FirebaseClient() *auth.Client {
	app := FirebaseApp()
	client, err := app.Auth(context.Background())
	if err != nil {
		log.Fatalf("error getting Auth client: %v\n", err)
	}
	return client
}
