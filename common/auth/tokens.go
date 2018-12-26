package auth

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"strings"
)

func getJwtToken(context *gin.Context) (*jwt.Token, error) {
	bearerToken := getBearerToken(context)

	jwtToken, parseError := jwt.Parse(bearerToken, getSigningKey)

	if parseError != nil {
		return nil, fmt.Errorf("JWT token could not be parsed", parseError)
	}

	if !jwtToken.Valid {
		return nil, errors.New("Parsed JWT is not valid")
	}

	return jwtToken, nil
}

func getBearerToken(context *gin.Context) string {
	authHeader := context.GetHeader("Authorization")
	token := strings.Replace(authHeader, "Bearer ", "", 1)
	return token
}

func getSigningKey(token *jwt.Token) (interface{}, error) {
	isValidAlgo := isSignedWithRsa256(token)

	if !isValidAlgo {
		return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
	}

	kid := token.Header["kid"].(string)

	return getRsaPublicKeyFromJwks(kid)
}

func isSignedWithRsa256(token *jwt.Token) bool {
	_, isRsa := token.Method.(*jwt.SigningMethodRSA)
	return isRsa
}
