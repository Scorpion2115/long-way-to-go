package dbs

import (
	"database/sql"
	//go get -u github.com/go-sql-driver/mysql
	_ "github.com/go-sql-driver/mysql" //we import supported libraries for database/sql
)

func Exec(db *sql.DB) error {
	// drop the table after the testing.
	defer db.Exec("DROP TABLE example")
	if err := Create(db); err != nil {
		return err
	}
	if err := Query(db, "Kevin Garnett"); err != nil {
		return err
	}
	return nil
}
