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

	stmt, _ := db.Prepare("update people set name=?, address=? where id=?")
	defer stmt.Close()

	r, _ := stmt.Exec("merry", "shandong", 3)
	count, _ := r.RowsAffected()
	if count > 0 {
		fmt.Println("ok")
	} else {
		fmt.Println("fail")
	}

}
