package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type OAuthResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    string `json:"expires_in"`
}

type UserClaim struct {
	jwt.RegisteredClaims
	Name string
}

func main() {
	router := gin.Default()
	router.GET("/order", Authenticate)
	router.Run()
}

func Authenticate(context *gin.Context) {
	var LOCAL_AUTH_SERVICE_URL = "http://localhost:4567"
	authServiceURL, exists := os.LookupEnv("AUTH_SERVICE_URL")
	if !exists {
		authServiceURL = LOCAL_AUTH_SERVICE_URL
	}

	var message string
	resp, err := http.Get(authServiceURL + "/signin")
	if err != nil {
		fmt.Println("[OrderService] Cannot contact Authentication service: " + authServiceURL)
		// os.Exit(0)
		message = "Cannot contact Authentication service"
	} else {
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)

		var tokenResponse OAuthResponse
		json.Unmarshal(body, &tokenResponse)

		// Parse JWT
		var userClaim UserClaim
		jwt.ParseWithClaims(tokenResponse.AccessToken, &userClaim, nil)
		message = fmt.Sprintf("Order placed for user %s (username: %s) | on %s", userClaim.Name, userClaim.Subject, userClaim.IssuedAt)
	}

	context.JSON(http.StatusOK, gin.H{"message": message})
}
