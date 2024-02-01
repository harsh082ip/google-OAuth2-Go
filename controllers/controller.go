package controllers

import (
	"context"
	"fmt"
	"io"
	"net/http"

	// "net/url"

	"github.com/harsh082ip/google-OAuth2-Go/config"
)

func GoogleLogin(res http.ResponseWriter, req *http.Request) {

	googleConfig := config.SetupConfig()
	url := googleConfig.AuthCodeURL("randomstate")

	// redirecting to login page
	http.Redirect(res, req, url, http.StatusSeeOther)
}

func GoogleCallback(res http.ResponseWriter, req *http.Request) {

	state := req.URL.Query()["state"][0]
	if state != "randomstate" {
		fmt.Fprintf(res, "status does not match")
	}

	code := req.URL.Query()["code"][0]

	gooleConfig := config.SetupConfig()

	token, err := gooleConfig.Exchange(context.Background(), code)
	if err != nil {
		fmt.Fprintf(res, "Code-Token Exchange Failed")
	}
	// fmt.Println(state)
	// fmt.Println()
	// fmt.Println(code)
	// fmt.Println()
	// fmt.Println(token)
	// fmt.Println()

	fmt.Println("---------------------------------------------")
	fmt.Println("Token Type: ", token.TokenType)
	fmt.Println("Refresh Token: ", token.RefreshToken)
	fmt.Println("Token Expiry: ", token.Expiry)
	fmt.Println("Access Token: ", token.AccessToken)
	resp, err := http.Get("https://www.googleapis.com/oauth2/v1/userinfo?alt=json&access_token=" + token.AccessToken)
	if err != nil {
		fmt.Fprintf(res, "User Data fetched failed")
	}

	userData, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintf(res, "Json Parsing Failed")
	}

	fmt.Fprintf(res, string(userData))
}
