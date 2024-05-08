package main

import (
	"context"
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	firebase "firebase.google.com/go"

	"github.com/ihummer/go-firebase/handlers"
	"google.golang.org/api/option"
)

func main() {
	// Initialize firebase
	opt := option.WithCredentialsFile("firebase-service-account.json")

	app, err := firebase.NewApp(context.Background(), nil, opt)

	if err != nil {
		log.Fatal(err)
	}

	// Initialize Firestore database
	db, err := app.Firestore(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Creating Echo instance
	e := echo.New()

	// Register middleware for error handling
	e.HTTPErrorHandler = handlers.ErrorHandler
	e.Use(middleware.Recover())

	// Creating petition handlers
	handlers.RegisterHandlers(e, db)

	// New http server
	fmt.Println("Starting server on port 8080...")
	log.Fatal(e.Start(":8080"))
}
