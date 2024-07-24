package handlers

// to be used for the auth logic
import (
	"log"
	"os"
	"github.com/zmb3/spotify"
	"github.com/joho/godotenv"
)

//exportable
var (
	Auth spotify.Authenticator
	State string
)
// main func to read in envir variables and set up
func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	clientID := os.Getenv("SPOTIFY_ID")
	clientSecret := os.Getenv("SPOTIFY_SECRET")
	RedirectUrl := os.Getenv("REDIRECT_URL")

	if clientID == "" || clientSecret == "" || RedirectUrl == "" {
        log.Fatal("environment variables are not set")
    }


	// we need to get a token
	Auth = spotify.NewAuthenticator(redirectURL, spotify.ScopePlaylistModifyPublic, spotify.ScopePlaylistModifyPrivate, streaming, user-read-email,user-read-private)
	Auth.SetAuthInfo(clientID, clientSecret)

	State = generateState()
}

func generateState() string {
    bytes := make([]byte, 16)
    _, err := rand.Read(bytes)
    if err != nil {
        log.Fatal(err)
    }
    return base64.URLEncoding.EncodeToString(bytes)
}