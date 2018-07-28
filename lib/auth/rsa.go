package auth

import (
	"crypto/rsa"
	"math/big"
	"bytes"
	"encoding/binary"
	"strings"
	"encoding/base64"
)

func createRsaPublicKey(jwk JWK) (*rsa.PublicKey, error){
	modulus, modulusError := getModulus(jwk.N)

	if(modulusError != nil){
		return nil, modulusError
	}

	exponent, exponentError := getExponent(jwk.E)

	if(exponentError != nil){
		return nil, exponentError
	}

	publicKey := rsa.PublicKey{
		N: modulus,
		E: exponent,
	}

	return &publicKey, nil
}

func getModulus(base64Modulus string)(*big.Int, error){
	decodedModulus, decodeError := safeBase64Decode(base64Modulus)

	if decodeError != nil {
		return nil, decodeError
	}

	modulus := big.NewInt(0)
	modulus.SetBytes(decodedModulus)

	return modulus, nil;
}

func getExponent(base64Exponent string)(int, error){
	decE, err := safeBase64Decode(base64Exponent)
	if err != nil {
		return -1, err
	}

	var eBytes []byte
	if len(decE) < 8 {
		eBytes = make([]byte, 8-len(decE), 8)
		eBytes = append(eBytes, decE...)
	} else {
		eBytes = decE
	}

	eReader := bytes.NewReader(eBytes)

	var exponent uint64

	err = binary.Read(eReader, binary.BigEndian, &exponent)

	if err != nil {
		return -1, err
	}

	return int(exponent), nil
}

func safeBase64Decode(str string) ([]byte, error) {
	lenMod4 := len(str) % 4
	if lenMod4 > 0 {
		str = str + strings.Repeat("=", 4-lenMod4)
	}

	return base64.URLEncoding.DecodeString(str)
}
