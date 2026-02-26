package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dsn := "root:root@tcp(127.0.0.1:3306)/test_table"

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("open db:", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal("ping db:", err)
	}

	rows, err := db.Query("SELECT id, name FROM test_table")
	if err != nil {
		log.Fatal("query:", err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string

		if err := rows.Scan(&id, &name); err != nil {
			log.Fatal("scan row:", err)
		}

		fmt.Printf("id=%d name=%s\n", id, name)
	}

	if err := rows.Err(); err != nil {
		log.Fatal("rows error:", err)
	}
}
