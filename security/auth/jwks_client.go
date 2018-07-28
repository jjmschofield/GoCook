package auth

import (
	"errors"
	"log"
	"github.com/jjmschofield/GoCook/net/jsonhttp"
	"fmt"
	"crypto/rsa"
)

const endpointUrl = "https://jjmschofield.eu.auth0.com/.well-known/jwks.json";

var keyCache []JWK

func getRsaPublicKey(kid string) (*rsa.PublicKey, error){
	jwk, jwkError := getJwk(kid);

	if(jwkError != nil){
		return nil, fmt.Errorf("JWK with kid matching %v is not available", kid)
	}

	publicKey, publicKeyError := createRsaPublicKey(jwk)

	if(publicKeyError != nil){
		return nil, publicKeyError
	}

	return publicKey, nil
}

func getJwk(kid string) (JWK, error) {

	for i := 0; i < 3; i++ {
		jwk, inCache := getJwkFromCache(kid)

		if(inCache) {
			return jwk, nil
		} else {
			log.Printf("Key %v not found in cache, refreshing key cache from JWKS endpoint", kid)
			syncJwksCache()
		}
	}

	return JWK{}, errors.New("JWK is not available")
}

func getJwkFromCache(kid string) (JWK, bool) {

	for i := range keyCache {
		if keyCache[i].Kid == kid {
			return keyCache[i], true
		}
	}

	return JWK{}, false;
}

func syncJwksCache() error{
	keys, err := getJwksFromEndpoint()

	if(err != nil){
		return err
	}

	keyCache = keys

	return nil
}

func getJwksFromEndpoint() ([]JWK, error){
	var jwks struct {
		Keys []JWK `json:"keys"`
	}

	error := jsonhttp.Get(endpointUrl, &jwks)

	return jwks.Keys, error
}
