package model

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var err error
var db *sql.DB

func init() {
	db, err = sql.Open("mysql", "homestead:secret@tcp([192.168.11.11]:3306)/aricto")
	if err != nil {
		panic(err.Error())
		fmt.Println(err)
	}
	// Test the connection to the database
	err = db.Ping()
	if err != nil {
		panic(err.Error())
		fmt.Println(err)
	}
}
