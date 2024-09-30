package utils

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var (
	// Secret key to sign JWT-tokens
	SecretKey string
)

// Generating new token with user info
//
// userInfo = { "id": int, "name": string }
func CreateToken(userInfo map[string]interface{}) (string, error) {
	// Create new token
	token := jwt.New(jwt.SigningMethodHS256)

	// Get claims of this token (payload)
	claims := token.Claims.(jwt.MapClaims)

	// Add some info to claims
	claims["name"] = userInfo["name"]
	claims["id"] = userInfo["id"]
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	// Get string token that will be passed to user
	// We sign this token with SecretKey
	tokenString, err := token.SignedString([]byte(SecretKey))
	if err != nil {
		return "", fmt.Errorf("error generate token: " + err.Error())
	}

	return tokenString, nil
}

// # Returns (claims, nil) if token is valid and all good
//
// # Returns (nil, error) if token is invalid for some reason
//
// Claims consists of name, id, exp(expiration time)
func ValidateToken(tokenSting string) (jwt.MapClaims, error) {
	// Parsing string version of token to *jwt.Token
	token, err := jwt.Parse(
		// First arg -> string to parse
		tokenSting,

		// Second arg -> function that check hash method of token
		// Return Secret key with what token string gonna be checked
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("wrong hash method: %v", token.Header["alg"])
			}

			return []byte(SecretKey), nil
		})

	if err != nil {
		return nil, fmt.Errorf("error in token: %v", err.Error())
	}

	// Check token expiration time and if it is valid
	// Get claims from parsed token
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		expirationTime := time.Unix(int64(claims["exp"].(float64)), 0)

		if expirationTime.Before(time.Now()) {
			return nil, fmt.Errorf("token has expired")
		}

		return claims, nil
	} else {
		return nil, fmt.Errorf("invalid token")
	}
}
