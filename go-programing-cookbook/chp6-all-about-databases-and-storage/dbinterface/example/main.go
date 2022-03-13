package main

import (
	"github.com/Scorpion2115/long-way-to-go/go-programing-cookbook/chp6-all-about-databases-and-storage/dbs"
	"dbinterface"
	//go get -u github.com/go-sql-driver/mysql
	_ "github.com/go-sql-driver/mysql" //we import supported libraries for database/sql
)

func main() {
	_, err_conn := dbs.TestConnection()
	if err_conn != nil {
		panic(err_conn)
	}
	tx,err :db.Begin()
	if err != nil{
		panic(err)
	}
	// this won't do anything if commit is successful
	defer tx.Rollback()
	if err : dbindbinterface.Exec(tx); err != nil {
		panic(err)
	}
	if err := tx.Commit(); err != nil {
		panic(err)
	}

}
