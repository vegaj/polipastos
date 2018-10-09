package auth

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"log"
)

//Sign is a part of the signature / validation service
//Uses HMAC SHA256
func Sign(message string, secret []byte) ([]byte, error) {

	h := hmac.New(sha256.New, secret)
	if _, err := h.Write([]byte(message)); err != nil {
		return nil, err
	}
	return h.Sum(nil), nil
}

//Verify checks that the message has been generated with the secret provided
func Verify(message, signature string, secret []byte) bool {

	var err error
	var generated []byte

	if generated, err = Sign(message, secret); err == nil && bytes.Compare(generated, []byte(signature)) == 0 {
		return true
	}

	return false
}

func toBase64(message []byte) string {
	//strings.TrimRight(base64.URLEncoding.EncodeToString(message), "=")
	return string(base64.URLEncoding.EncodeToString(message))
}

func fromBase64String(message string) []byte {
	dec, err := base64.URLEncoding.DecodeString(message)
	if err != nil {
		log.Println("Error", err)
		return nil
	}
	return dec
}
