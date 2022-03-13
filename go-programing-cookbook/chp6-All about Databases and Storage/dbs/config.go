package dbs

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	//go get -u github.com/go-sql-driver/mysql
	// this is the driver for you to connect various database
	"github.com/go-sql-driver/mysql"
)

// Example hold the results of our queries
type Example struct {
	Name    string
	Created *time.Time
}

// Testing Connection
func TestConnection() (*sql.DB, error) {
	cfg := mysql.Config{
		User:                 "evan",
		Passwd:               "tpg12345",
		Net:                  "tcp",
		Addr:                 "10.224.44.127:30010",
		DBName:               "demo",
		AllowNativePasswords: true,
	}

	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		return nil, err
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Printf("Testing Connection.....Successfully")
	return db, nil
}

// Setup configures and returns the databases
func Setup() (*sql.DB, error) {
	//db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(10.224.44.127:30010)/demo?parseTime=true", os.Getenv("MYSQLUSERNAME"), os.Getenv("MYSQLPASSWORD")))
	db, err := sql.Open("mysql", "evan:tpg12345@tcp(10.224.44.127:30010)/demo?parseTime=true")
	if err != nil {
		return nil, err
	}
	return db, nil
}
