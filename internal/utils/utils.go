package internal

import (
	"crypto/md5"
	"strings"
	"encoding/base64"
)

// Creates a unique hash, but isn't really unique b/c of truncation ... 
func genUniqueHash(str string) string {
	md5Hasher := md5.New()
	md5Hasher.Write([]byte(str))
	hash :=	md5Hasher.Sum(nil)
	return base64.StdEncoding.EncodeToString(hash[:8])
}

// Public function
func MakeHash (str string) string {
	return strings.Trim(genUniqueHash(str), "==")
}
