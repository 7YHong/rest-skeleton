package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

var HmacKey string

func HmacSha256(key, data string) string {
	hash := hmac.New(sha256.New, []byte(key)) //创建对应的sha256哈希加密算法
	hash.Write([]byte(data))
	return hex.EncodeToString(hash.Sum(nil))
}
