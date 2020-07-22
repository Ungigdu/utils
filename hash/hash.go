package hash

import (
	"crypto/md5"
	"encoding/hex"
)

func HashMD5(bytes []byte) string{
	hash := md5.New()
	hash.Write(bytes)
	return hex.EncodeToString(hash.Sum(nil))
}