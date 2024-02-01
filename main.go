package main

import (
	"fmt"
	"net/http"

	"github.com/harsh082ip/google-OAuth2-Go/controllers"
)

const (
	WEBPORT = ":8001"
)

func main() {

	http.HandleFunc("/", controllers.GoogleLogin)
	http.HandleFunc("/google/callback", controllers.GoogleCallback)

	fmt.Printf("Server Running on %v...", WEBPORT)
	http.ListenAndServe(WEBPORT, nil)
}
