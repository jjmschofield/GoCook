package auth

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const expectedIssuer = "https://jjmschofield.eu.auth0.com/"
const expectedAudience = "https://go-cook"

func hasValidStandardClaims(token *jwt.Token) (bool, error){
	claims := token.Claims.(jwt.MapClaims)

	invalidExp, expError := hasExpired(claims["exp"].(float64))
	if(invalidExp){
		return false, expError
	}

	invalidAudience, audError := hasInvalidAudience(claims["aud"].(string))
	if(invalidAudience){
		return false, audError
	}

	invalidIssuer, issError := hasInvalidIssuer(claims["iss"].(string))
	if(invalidIssuer){
		return false, issError
	}

	return true, nil
}

func hasExpired(exp float64) (bool, error){
	expiryTime := time.Unix(int64(exp), 0)

	if(expiryTime.Before(time.Now())){
		return true, fmt.Errorf("Token expired at %v", expiryTime)
	}

	return false, nil
}

func hasInvalidAudience(aud string)(bool, error){
	if(aud != expectedAudience){
		return true, fmt.Errorf("Token audience %v is not valid", aud)
	}

	return false, nil
}

func hasInvalidIssuer(iss string)(bool, error){
	if(iss != expectedIssuer){
		return true, fmt.Errorf("Token audience %v is not valid", iss)
	}

	return false, nil
}
