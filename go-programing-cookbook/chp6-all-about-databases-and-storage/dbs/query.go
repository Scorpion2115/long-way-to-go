package dbs

import (
	"database/sql"
	"fmt"

	//go get -u github.com/go-sql-driver/mysql
	_ "github.com/go-sql-driver/mysql" //we import supported libraries for database/sql
)

// Query grabs a new connection
// creates tables, and later drops them
// and issues some queries
func Query(db *sql.DB, name string) error {
	//name := "Kevin Garnett"
	rows, err := db.Query("SELECT name, created FROM example")
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		var e Example
		if err := rows.Scan(&e.Name, &e.Created); err != nil {
			return err
		}
		fmt.Printf("Results:\n\tName: %s\n\tCreated: %v\n", e.Name, e.Created)
	}
	return rows.Err()
}
