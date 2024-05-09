package pkg

import (
	"crypto/md5"
	"encoding/hex"
)

const secret = "dsj.com"

func encryptPassword(password string) string {
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum([]byte(password)))
}
