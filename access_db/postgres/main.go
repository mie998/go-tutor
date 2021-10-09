package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type EMPLOYEE struct {
	ID     string
	NUMBER string
}

// reference: https://qiita.com/hiro9/items/e6e41ec822a7077c3568
func main() {
	db, err := sql.Open("postgres", "host=127.0.0.1 port=5555 user=root password=password dbname=testdb sslmode=disable")
	defer db.Close()
	if err != nil {
		fmt.Println(err)
	}

	// INSERT
	var empID string
	id := 3
	number := 4444
	err = db.QueryRow("INSERT INTO employee(emp_id, emp_number) VALUES($1,$2) RETURNING emp_id", id, number).Scan(&empID)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(empID)

	// SELECT
	rows, err := db.Query("SELECT * FROM employee")
	if err != nil {
		fmt.Println(err)
	}

	var es []EMPLOYEE
	for rows.Next() {
		var e EMPLOYEE
		rows.Scan(&e.ID, &e.NUMBER)
		es = append(es, e)
	}
	fmt.Println(es)
}
