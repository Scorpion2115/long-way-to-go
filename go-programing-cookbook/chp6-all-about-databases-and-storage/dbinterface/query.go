package dbinterface

import (
	//go get -u github.com/go-sql-driver/mysql
	"fmt"

	"github.com/Scorpion2115/long-way-to-go/go-programing-cookbook/chp6-all-about-databases-and-storage/dbs"

	_ "github.com/go-sql-driver/mysql" //we import supported libraries for database/sql
)

func Query(db DB) error {
	name := "Kevin Garnett"
	rows, err := db.Query("SELECT name, created FROM example where name=?", name)
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		var e dbs.Example
		if err := rows.Scan(&e.Name, &e.Created); err != nil {
			return err
		}
		fmt.Printf("Results:\n\tName: %s\n\tCreated: %v\n", e.Name, e.Created)
	}
	return rows.Err()
}
