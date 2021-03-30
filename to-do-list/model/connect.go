package model

import (
	"database/sql"
	"fmt"
	"log"
)

var con *sql.DB     // all db operations will have access to this connection only
func Connect() *sql.DB {

	//db, err := sql.Open("mysql", "root:root1234@(tcp:localhost:3306)")

	dsn := "root:root1234@/test"
	db, err := sql.Open("mysql", dsn)
	if err != nil{
		log.Fatal(err)
	}else{
		fmt.Println("Database connnected...")
	}

	con = db
	return db
}