package internal

import (
	"crypto/md5"
	"strings"
	"encoding/base64"
)

// Hashing scaffolding
func genUniqueHash(str string) string {
	md5Hasher := md5.New()
	md5Hasher.Write([]byte(str))
	hash :=	md5Hasher.Sum(nil)
	return base64.StdEncoding.EncodeToString(hash[:5])

}

func MakeHash (str string) string {
	return strings.Trim(genUniqueHash(str), "==")
}

/* 
func Nothing () int {
	return 3434;
}
*/
