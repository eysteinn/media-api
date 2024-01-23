package database

import (
	"database/sql"
	_ "embed"
	"log"
	"os"
	"path"

	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB = nil
var dbfile = "photos.sqlite"

//go:embed photos_create.sql
var sql_photos_create string

//go:embed faces_create.sql
var sql_faces_create string

func GetDB() *sql.DB {
	if db == nil {
		Init()
	}
	return db
}
func tableExists(db *sql.DB, tableName string) (bool, error) {
	var exists bool
	query := "SELECT name FROM sqlite_master WHERE type='table' AND name=?"
	err := db.QueryRow(query, tableName).Scan(&exists)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	return exists, nil
}
func Init() error {
	if db == nil {
		fileexists := fileExists(dbfile)
		tmpdb, err := sql.Open("sqlite3", dbfile)
		if err != nil {
			return err
		}
		db = tmpdb
		/*sql_photos_create := `
		CREATE TABLE IF NOT EXISTS files (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			filename TEXT NOT NULL,
			unixtime INTEGER,
			sha256 TEXT NOT NULL,
			INDEX idx_unixtime (unixtime)
		);
		`*/
		// Execute the SQL statement to create the table
		if fileexists == false {
			_, err = db.Exec(sql_photos_create)
			if err != nil {
				return err
			}
		}

		exists, err := tableExists(db, "faces")
		if err != nil {
			return err
		}
		if !exists {
			log.Println("Creating faces table")
			_, err = db.Exec(sql_faces_create)
			if err != nil {
				return err
			}
		}

	}
	return nil
}

var uuidV4 *uuid.UUID = nil

func InsertPhoto(filepath string, unixtime int32, sha256Hash string) error {
	/*if uuidV4 == nil {
		tmp := uuid.New()
		uuidV4 = &tmp
	}*/

	filename := path.Base(filepath)
	// Insert data into the table using Unix time
	insertSQL := `
	INSERT INTO files (filename, filepath, unixtime, sha256, uuid) VALUES (?, ?, ?, ?, ?)
	`

	// Sample data to be inserted
	/*filename := "example.txt"
	unixtime := time.Now().Unix()   // Current Unix time
	sha256Hash := "a1b2c3d4e5f6..." // Replace with the actual SHA-256 hash
	*/
	// Execute the SQL statement to insert data
	err := Init()
	if err != nil {
		return err
	}
	db := GetDB()
	_, err = db.Exec(insertSQL, filename, filepath, unixtime, sha256Hash, uuid.NewString())
	return err
}

/*func CreateTable() {
	// Open the database file or create it if it doesn't exist
	db, err := sql.Open("sqlite3", "test.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create the table
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		age INTEGER
	);
	`

	// Execute the SQL statement to create the table
	_, err = db.Exec(createTableSQL)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Table 'users' created successfully")
}*/

// fileExists checks if a file exists at the specified path
func fileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		return false
	}
	return true
}
