package middleware

import (
	"errors"
	"github.com/corysakti/cats-social-go/helper"
	"github.com/corysakti/cats-social-go/model/web/response"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"strings"
)

type AuthMiddleware struct {
	Handler http.Handler
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddleware {
	return &AuthMiddleware{Handler: handler}
}

func (middleware *AuthMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	// Define whitelisted URLs
	whitelistedURLs := []string{
		"/v1/user/register",
		"/v1/user/login",
	}

	// Check if the requested URL is whitelisted
	for _, url := range whitelistedURLs {
		if request.URL.Path == url {
			// If the URL is whitelisted, call the next handler without JWT token validation
			middleware.Handler.ServeHTTP(writer, request)
			return
		}
	}

	// Extract the JWT token from the Authorization header
	authHeader := request.Header.Get("Authorization")
	if authHeader == "" {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusUnauthorized)
		responseTemplate := response.ResponseTemplate{
			Code:   http.StatusUnauthorized,
			Status: "UNAUTHORIZED",
		}
		helper.WriteToResponseBody(writer, responseTemplate)
		return
	}
	tokenString := strings.Replace(authHeader, "Bearer ", "", 1)

	// Parse and validate the JWT token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Check the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid token")
		}
		// Provide the key used for signing
		return []byte("your-secret-key"), nil
	})
	if err != nil {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusUnauthorized)
		responseTemplate := response.ResponseTemplate{
			Code:   http.StatusUnauthorized,
			Status: "UNAUTHORIZED",
		}
		helper.WriteToResponseBody(writer, responseTemplate)
		return
	}

	// If the token is valid, set the user information in the request context
	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Assuming your JWT token has a userID claim
		// Fetch user details from the database or validate user existence based on userID
		// Here you can also set the user details in the request context if needed
		// For simplicity, let's assume the user is valid
		// Check user existence or fetch user details from DB using middleware.Db
	} else {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusUnauthorized)
		responseTemplate := response.ResponseTemplate{
			Code:   http.StatusUnauthorized,
			Status: "UNAUTHORIZED",
		}
		helper.WriteToResponseBody(writer, responseTemplate)
		return
	}

	// If everything is fine, call the next handler
	middleware.Handler.ServeHTTP(writer, request)
}
