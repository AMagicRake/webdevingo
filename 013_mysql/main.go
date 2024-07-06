package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	dbUser := os.Args[1]
	dbPass := os.Args[2]
	dbHost := os.Args[3]
	dbPort := os.Args[4]
	dbName := os.Args[5]
	//user:password@tcp(db:port)/dbName?charset=utf8
	connString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", dbUser, dbPass, dbHost, dbPort, dbName)
	fmt.Println(connString)
	conn, err := sql.Open("mysql", connString)
	if err != nil {
		log.Fatalln(err)
	}
	conn.Close()
}
