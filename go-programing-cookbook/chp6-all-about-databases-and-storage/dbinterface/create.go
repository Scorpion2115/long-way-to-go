package dbinterface

import (
	//go get -u github.com/go-sql-driver/mysql
	"fmt"

	_ "github.com/go-sql-driver/mysql" //we import supported libraries for database/sql
)

func Create(db DB) error {
	if _, err := db.Exec("CREATE TABLE example (name VARCHAR(20), created DATETIME)"); err != nil {
		return err
	}

	if _, err := db.Exec(`INSERT INTO example (name, created) values ("Kevin Garnett", NOW())`); err != nil {
		return err
	}
	fmt.Printf("\nTable Created\n")
	return nil
}
