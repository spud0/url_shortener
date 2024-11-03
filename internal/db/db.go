package internal

import (
	"database/sql"
	"github.com/lib/pq"
)

// database scaffolding
// Need to figure out the DB schema
type DB struct {
	Conn *sql.DB
}

func (db *DB) SaveURL (url string) error {
	_, err := db.Conn.Exec("...")
	return err	
}

func (db *DB) GetURL (hashedKey string) (string, error) {
	var longURL string
	err := db.Conn.QueryRow("...", hashedKey).Scan(&longURL)
	return longURL, err	
}
	
