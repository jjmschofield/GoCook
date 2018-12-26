package auth

import (
	"bytes"
	"crypto/rsa"
	"encoding/base64"
	"encoding/binary"
	"math/big"
	"strings"
)

func createRsaPublicKey(jwk jsonWebKey) (*rsa.PublicKey, error) {
	modulus, modulusError := getModulusBigInt(jwk.N)

	if modulusError != nil {
		return nil, modulusError
	}

	exponent, exponentError := getExponentInt(jwk.E)

	if exponentError != nil {
		return nil, exponentError
	}

	publicKey := rsa.PublicKey{
		N: modulus,
		E: exponent,
	}

	return &publicKey, nil
}

func getModulusBigInt(base64Modulus string) (*big.Int, error) {
	decodedModulus, decodeError := safeBase64Decode(base64Modulus)

	if decodeError != nil {
		return nil, decodeError
	}

	modulus := big.NewInt(0)
	modulus.SetBytes(decodedModulus)

	return modulus, nil
}

func getExponentInt(base64Exponent string) (int, error) {
	decodedExponent, decodeError := safeBase64Decode(base64Exponent)

	if decodeError != nil {
		return -1, decodeError
	}

	paddedExponent := getPaddedExponent(decodedExponent)

	binaryReader := bytes.NewReader(paddedExponent)

	var exponent uint64

	binaryReadError := binary.Read(binaryReader, binary.BigEndian, &exponent)

	if binaryReadError != nil {
		return -1, binaryReadError
	}

	return int(exponent), nil
}

func getPaddedExponent(exponent []byte) []byte {
	var paddedExponent []byte

	if len(exponent) < 8 {
		paddedExponent = make([]byte, 8-len(exponent), 8)
		paddedExponent = append(paddedExponent, exponent...)
	} else {
		paddedExponent = exponent
	}

	return paddedExponent
}

func safeBase64Decode(str string) ([]byte, error) {
	lenMod4 := len(str) % 4
	if lenMod4 > 0 {
		str = str + strings.Repeat("=", 4-lenMod4)
	}

	return base64.URLEncoding.DecodeString(str)
}
