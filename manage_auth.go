package main

import (
	"log"
	"os"

	//"github.com/zmb3/spotify"
	//"golang.org/x/oauth2"
	//"golang.org/x/oauth2/clientcredentials"
	"github.com/joho/godotenv"
)

// main func to read in envir variables and set up
func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	clientID := os.Getenv("SPOTIFY_ID")
	clientSecret := os.Getenv("SPOTIFY_SECRET")

	// we need to get a token

}
