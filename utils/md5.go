package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func MD5Encode(code string) string {
	h := md5.New()
	h.Write([]byte(code))
	return hex.EncodeToString(h.Sum(nil))
}

func MakePassword(plainpwd, salt string) string {
	return MD5Encode(plainpwd + salt)
}

func ValidPassword(plainpwd, salt, password string) bool {
	return MakePassword(plainpwd, salt) == password
}
