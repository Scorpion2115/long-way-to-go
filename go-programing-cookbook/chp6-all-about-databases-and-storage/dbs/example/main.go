package main

import (
	//"dbs"
	"github.com/Scorpion2115/long-way-to-go/go-programing-cookbook/chp6-all-about-databases-and-storage/dbs"

	//go get -u github.com/go-sql-driver/mysql
	_ "github.com/go-sql-driver/mysql" //we import supported libraries for database/sql
)

func main() {
	_, err_conn := dbs.TestConnection()
	if err_conn != nil {
		panic(err_conn)
	}

	db, err := dbs.Setup()
	if err != nil {
		panic(err)
	}

	if err := dbs.Exec(db); err != nil {
		panic(err)
	}
}
