package internal

import (
	"database/sql"
	"github.com/lib/pq"
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


// Caching scaffolding
type Cache struct {
	Client *redis.Client
}

func (c *Cache) CacheURL () {

}

func (c *Cache) GetURLFromCache () {

}


// database scaffolding
type DB struct {
	Conn *sql.DB
}

func (db *DB) SaveURL (url string) (bool, error) {
	
}

func (db *DB) GetURL (url string) (string, error) {

}
	
func Nothing () int {
	return 3434;
}
