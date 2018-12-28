package auth

import (
	"crypto/rsa"
	"errors"
	"fmt"
	"github.com/jjmschofield/GoCook/common/jsonhttp"
	"github.com/jjmschofield/GoCook/common/logger"
	"github.com/spf13/viper"
)

type jsonWebKey struct {
	Alg string   `json:"alg"`
	Kty string   `json:"kty"`
	Use string   `json:"use"`
	X5C []string `json:"x5c"`
	N   string   `json:"n"`
	E   string   `json:"e"`
	Kid string   `json:"kid"`
	X5T string   `json:"x5t"`
}

var jwksCache []jsonWebKey

func getRsaPublicKeyFromJwks(kid string) (*rsa.PublicKey, error) {
	jwk, jwkError := getJwk(kid)

	if jwkError != nil {
		return nil, fmt.Errorf("jsonWebKey with kid matching %v is not available", kid)
	}

	publicKey, publicKeyError := createRsaPublicKey(jwk)

	if publicKeyError != nil {
		return nil, publicKeyError
	}

	return publicKey, nil
}

func getJwk(kid string) (jsonWebKey, error) {

	for i := 0; i < 3; i++ {
		jwk, inCache := getJwkFromCache(kid)

		if inCache {
			return jwk, nil
		} else {
			logger.Debug(fmt.Sprintf("Key %v not found in cache, refreshing key cache from JWKS endpoint", kid))
			syncJwksCache()
		}
	}

	return jsonWebKey{}, errors.New("jsonWebKey is not available")
}

func getJwkFromCache(kid string) (jsonWebKey, bool) {

	for i := range jwksCache {
		if jwksCache[i].Kid == kid {
			return jwksCache[i], true
		}
	}

	return jsonWebKey{}, false
}

func syncJwksCache() error {
	keys, err := getJwksFromEndpoint()

	if err != nil {
		return err
	}

	jwksCache = keys

	return nil
}

func getJwksFromEndpoint() ([]jsonWebKey, error) {
	var jwks struct {
		Keys []jsonWebKey `json:"keys"`
	}

	endpointUrl := viper.GetString("AUTH_JWKS_ENDPOINT")

	requestError := jsonhttp.Get(endpointUrl, &jwks)

	return jwks.Keys, requestError
}
