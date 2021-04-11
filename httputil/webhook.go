package httputil

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"os"
	"strings"
)

func VerifySignature(signature string, body []byte) bool {
	secret := []byte(os.Getenv("GITHUB_WEBHOOK_SECRET"))
	const signaturePrefix = "sha1="
	const signatureLength = 45
	if len(signature) != signatureLength || !strings.HasPrefix(signature, signaturePrefix) {
		return false
	}
	actual := make([]byte, 20)
	hex.Decode(actual, []byte(signature[5:]))
	generated := generateBodySignature(secret, body)
	return hmac.Equal(generated, actual)
}

func generateBodySignature(secret, body []byte) []byte {
	computed := hmac.New(sha1.New, secret)
	computed.Write(body)
	return []byte(computed.Sum(nil))
}
