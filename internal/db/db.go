package db 

import (
	"database/sql"
	_ "github.com/lib/pq"
	"os"
	"fmt"
)

// database scaffolding
// Need to figure out the DB schema
type DB struct {
	Conn *sql.DB
}

// Initializes a new connection to the database.
func NewDB ()  (*DB, error) {
	
	// Getting the environment variables
	port, dbName := os.Getenv("PORT"), os.Getenv("NAME")
	user, passwd:= os.Getenv("USER"), os.Getenv("PASS")
	hostName := os.Getenv("HOST")

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			hostName, port, user, passwd, dbName)

	conn, err := sql.Open("postgres", psqlInfo)
	if (err != nil) {
		return nil, err	
	}

	fmt.Println("connected to db")
	return &DB{Conn: conn}, nil
	
}

// Puts the URL into the DB.
func (db *DB) SaveURL (url string) error {
	_, err := db.Conn.Exec("...")
	return err	
}


// Gets the long URL from the DB.
func (db *DB) GetURL (hashedKey string) (string, error) {
	var longURL string
	err := db.Conn.QueryRow ("SELECT long_url FROM urls WHERE short_key = $1", hashedKey).Scan(&longURL)
	return longURL, err	
}
	
