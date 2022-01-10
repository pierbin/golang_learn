package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type user struct {
	id        int
	username  string
	password  string
	createdAt time.Time
}

// Create a new table
func createUsersTable(db *sql.DB) {
	query := `
		    CREATE TABLE users (
		        id INT AUTO_INCREMENT,
		        username TEXT NOT NULL,
		        password TEXT NOT NULL,
		        created_at DATETIME,
		        PRIMARY KEY (id)
		    );`

	if _, err := db.Exec(query); err != nil {
		// If the users table has been created, go ahead, don't show the error.
		if err.Error() != "Error 1050: Table 'users' already exists" {
			log.Fatal(err)
		}
	}
}

// Insert a new user
func insertRow(db *sql.DB) (int64, error) {
	username := "john doe"
	password := "secret"
	createdAt := time.Now()

	result, err := db.Exec(`INSERT INTO users (username, password, created_at) VALUES (?, ?, ?)`, username, password, createdAt)
	if err != nil {
		log.Fatal(err)
	}

	return result.LastInsertId()
}

// Query a single user
func getRowById(db *sql.DB, userId int64) {
	var (
		id        int
		username  string
		password  string
		createdAt time.Time
	)

	query := "SELECT id, username, password, created_at FROM users WHERE id = ?"
	if err := db.QueryRow(query, userId).Scan(&id, &username, &password, &createdAt); err != nil {
		log.Fatal(err)
	}

	fmt.Println(id, username, password, createdAt)
}

// Get all users
func getAllRows(db *sql.DB) {
	rows, err := db.Query(`SELECT id, username, password, created_at FROM users`)
	if err != nil {
		log.Fatal(err)
	}

	var users []user
	for rows.Next() {
		var u user

		err := rows.Scan(&u.id, &u.username, &u.password, &u.createdAt)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, u)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%#v\n", users)
}

// Delete user by userId
func delRowById(db *sql.DB, userId int64) {
	_, err := db.Exec(`DELETE FROM users WHERE id = ?`, userId)
	if err != nil {
		log.Fatal(err)
	}
}

// Delete all users
func delAllRows(db *sql.DB) {
	_, err := db.Exec(`DELETE FROM users`)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	// db, err := sql.Open("mysql", "username:password@(127.0.0.1:3306)/dbname?parseTime=true")
	db, err := sql.Open("mysql", "root:123456@(127.0.0.1:3306)/go_web?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}

	// Ping verifies a connection to the database is still alive, establishing a connection if necessary.
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	createUsersTable(db)
	id, err := insertRow(db)
	fmt.Printf("Last insert user id is %v \n", id)

	getRowById(db, id)
	delRowById(db, id)
	delAllRows(db)
	getAllRows(db)
}
