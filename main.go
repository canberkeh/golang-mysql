package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

type testusers struct {
	id    int
	name  string
	email string
}

func main() {
	// Capture connection properties.
	cfg := mysql.Config{
		User:   "root",        // Set Username Here
		Passwd: "pass", // Set Password Here
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "testdb", // Set Database Name Here
	}
	// Get a database handle.
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")

	values, err := testUserById(1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("User found: %v\n", values)

	userId, err := addUser(testusers{
		name:  "Test User",
		email: "testuser@mail.com"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("ID of added users: %v\n", userId)
}

// testusers by id queries for testusers that have the specified id.
func testUserById(id int) ([]testusers, error) {
	// An testusers slice to hold data from returned rows.
	var users []testusers

	rows, err := db.Query("SELECT * FROM testusers WHERE id = ?", id)
	if err != nil {
		return nil, fmt.Errorf("testUserById %q: %v", id, err)
	}
	defer rows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var user testusers
		if err := rows.Scan(&user.id, &user.name, &user.email); err != nil {
			return nil, fmt.Errorf("testUserById %q: %v", id, err)
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("testUserById %q: %v", id, err)
	}
	return users, nil
}

// addUser adds the specified user to the database,
// returning the user ID of the new entry
func addUser(user testusers) (int64, error) {
	result, err := db.Exec("INSERT INTO testusers (name, email) VALUES (?, ?)", user.name, user.email)
	if err != nil {
		return 0, fmt.Errorf("addUser: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("addUser: %v", err)
	}
	return id, nil
}
