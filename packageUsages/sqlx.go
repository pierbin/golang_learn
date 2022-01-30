package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
)

type User struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
	Email string `db:"email"`
}

// It is used only when a table does not exist.
var schema = "CREATE TABLE `users` (" +
	"`id` integer AUTO_INCREMENT NOT NULL PRIMARY KEY," +
	"`name` varchar(255) NOT NULL," +
	"`email` varchar(255) NOT NULL)"

func insertUser(db *sqlx.DB) int64 {
	res, err := db.Exec("INSERT INTO users (name, email) VALUES(\"Peter\", \"davy@gmail.com\")")
	if err != nil {
		panic(err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		panic(err)
	}

	return id
}

func getUserById(db *sqlx.DB, id int64) {
	var users User
	err := db.Get(&users, "select * from users where id=?", id)
	if err != nil {
		panic(err)
	}

	fmt.Println(users)
}

func updateUserById(db *sqlx.DB, id int64) {
	_, err := db.Exec("UPDATE users set email=\"Test@gmail.com\" where id=?", id)
	if err != nil {
		panic(err)
	}
}

func delUserById(db *sqlx.DB, id int64) {
	_, err := db.Exec("DELETE FROM users where id=?", id)
	if err != nil {
		panic(err)
	}
}

func getAllUsers(db *sqlx.DB) ([]User, error) {
	var users []User
	err := db.Select(&users, "SELECT * FROM users")

	return users, err
}

func main() {
	db, err := sqlx.Connect("mysql", "root:123456@(localhost:3306)/go_web")
	if err != nil {
		log.Fatalln(err)
	}

	// Only exec when the table needs to create. If the table has been created, it will have an error.
	// db.MustExec(schema)

	id := insertUser(db)
	fmt.Printf("Created user with id:%d\n", id)

	getUserById(db, id)
	updateUserById(db, id)
	delUserById(db, id-1)

	users, err := getAllUsers(db)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(users)
}
