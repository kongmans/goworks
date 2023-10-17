package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {

	db, err := sql.Open("mysql", "root:root123456@tcp(localhost:3306)/test")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	// Use connection pooling
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)

	//Use prepared statements
	stmt, err := db.Prepare("SELECT * FROM people")
	//stmt, err := db.Prepare("SELECT * FROM people WHERE id = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	// Use query caching
	rows, err := stmt.Query()
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var id int
		var name string
		var addr string
		rows.Scan(&id, &name, &addr)
		fmt.Println(id, name, addr)
	}

	defer rows.Close()
}
