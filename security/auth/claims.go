package auth

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
	"github.com/spf13/viper"
)

func hasValidClaims(token *jwt.Token) (bool, error){

	invalidExp, expError := hasExpired(token)
	if(invalidExp){
		return false, expError
	}

	invalidAudience, audError := hasInvalidAudience(token)
	if(invalidAudience){
		return false, audError
	}

	invalidIssuer, issError := hasInvalidIssuer(token)
	if(invalidIssuer){
		return false, issError
	}

	return true, nil
}

func hasExpired(token *jwt.Token) (bool, error){
	exp := token.Claims.(jwt.MapClaims)["exp"].(float64) // jwt-go is transforming this to a float64 undesirably

	expiryTime := time.Unix(int64(exp), 0)

	if(expiryTime.Before(time.Now())){
		return true, fmt.Errorf("Token expired at %v", expiryTime)
	}

	return false, nil
}

func hasInvalidAudience(token *jwt.Token)(bool, error){
	aud := token.Claims.(jwt.MapClaims)["aud"].(string)

	expectedAudience := viper.GetString("AUTH_AUDIENCE")

	if(aud != expectedAudience){
		return true, fmt.Errorf("Token audience %v is not valid", aud)
	}

	return false, nil
}

func hasInvalidIssuer(token *jwt.Token)(bool, error){
	iss := token.Claims.(jwt.MapClaims)["iss"].(string)

	expectedIssuer := viper.GetString("AUTH_ISSUER")

	if(iss != expectedIssuer){
		return true, fmt.Errorf("Token audience %v is not valid", iss)
	}

	return false, nil
}

