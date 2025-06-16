package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"errors"
)

func Sign(merchantId string, secretKey string) (string, error) {
	// 1. Validate key
	if secretKey == "" {
		return "", errors.New("APP_KEY 参数为空，请填写")
	}

	h := hmac.New(sha256.New, []byte(secretKey))
	h.Write([]byte(merchantId))
	signResult := hex.EncodeToString(h.Sum(nil))

	return signResult, nil
}

func Verify(merchantId string, secretKey string, signKey string) (bool, error) {
	// Check if signature exists in params
	currentSignature, _ := Sign(merchantId, secretKey)

	// Compare signatures
	return signKey == currentSignature, nil
}
