package auth

import (
	"errors"
	"log"
	"github.com/jjmschofield/GoCook/src/lib/jsonHttp"
)

const endpointUrl = "https://jjmschofield.eu.auth0.com/.well-known/jwks.json";

var keyCache []JWK

func GetSigningKey(kid string) (JWK, error) {

	for i := 0; i < 3; i++ {
		jwk, inCache := getKeyFromCache(kid)

		if(inCache) {
			return jwk, nil
		}else{
			log.Print("Key not found in cache, requesting keys from JWKS endpoint")
			updateSigningKeysCache()
		}
	}

	return JWK{}, errors.New("JWK is not available")
}

func getKeyFromCache(kid string) (JWK, bool) {

	for i := range keyCache {
		if keyCache[i].Kid == kid {
			return keyCache[i], true
		}
	}

	return JWK{}, false;
}

func updateSigningKeysCache() error{
	keys, err := GetJWKS()

	if(err != nil){
		return err
	}

	keyCache = keys

	return nil
}

func GetJWKS() ([]JWK, error){
	var jwks struct {
		Keys []JWK `json:"keys"`
	}

	error := jsonHttp.Get(endpointUrl, &jwks);

	return jwks.Keys, error
}
