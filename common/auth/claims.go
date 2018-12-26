package auth

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
	"time"
)

func GetClaim(token *jwt.Token, claim string) string {
	return token.Claims.(jwt.MapClaims)[claim].(string)
}

func hasValidClaims(token *jwt.Token) (bool, error) {

	invalidExp, expError := hasExpired(token)
	if invalidExp {
		return false, expError
	}

	invalidAudience, audError := hasInvalidAudience(token)
	if invalidAudience {
		return false, audError
	}

	invalidIssuer, issError := hasInvalidIssuer(token)
	if invalidIssuer {
		return false, issError
	}

	return true, nil
}

func hasExpired(token *jwt.Token) (bool, error) {
	exp := token.Claims.(jwt.MapClaims)["exp"].(float64) // jwt-go is transforming this to a float64 undesirably

	expiryTime := time.Unix(int64(exp), 0)

	if expiryTime.Before(time.Now()) {
		return true, fmt.Errorf("Token expired at %v", expiryTime)
	}

	return false, nil
}

func hasInvalidAudience(token *jwt.Token) (bool, error) {
	requiredAudience := viper.GetString("AUTH_AUDIENCE")

	gty := token.Claims.(jwt.MapClaims)["gty"]

	if gty == "client-credentials" {
		audience := token.Claims.(jwt.MapClaims)["aud"]

		if audience == requiredAudience {
			return false, nil
		}

		return true, fmt.Errorf("Token audience %v is not valid", audience)
	} else {
		audienceClaims := token.Claims.(jwt.MapClaims)["aud"].([]interface{}) // When the token is generated from the implicit flow it may have multiple audiences

		for _, claim := range audienceClaims {
			if claim == requiredAudience {
				return false, nil
			}
		}

		return true, fmt.Errorf("token audienceClaims %v is not valid", audienceClaims)
	}
}

func hasInvalidIssuer(token *jwt.Token) (bool, error) {
	iss := token.Claims.(jwt.MapClaims)["iss"].(string)

	expectedIssuer := viper.GetString("AUTH_ISSUER")

	if iss != expectedIssuer {
		return true, fmt.Errorf("Token audience %v is not valid", iss)
	}

	return false, nil
}
