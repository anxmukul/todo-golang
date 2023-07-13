package database

// package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

type Db struct {
	Db *sql.DB
}

var (
	host     = os.Getenv("host")
	port     = os.Getenv("port")
	user     = os.Getenv("user")
	password = os.Getenv("password")
	dbname   = os.Getenv("dbname")
)

func (d Db) ConnectToDb() *sql.DB {
	fmt.Println("Inside connectToDb() function")
	psqlconn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		// panic(err)
		fmt.Println("Db connectio failed")
		os.Exit(1)
	}
	return db
}

// func handleRowsReturedByQuery(rows *sql.Rows) int {
// 	var noOfRows = 0
// 	var TodoId int
// 	var TodoTitle string
// 	var TodoContent string
// 	for rows.Next() {
// 		err := rows.Scan(&TodoId, &TodoTitle, &TodoContent)
// 		if err != nil {
// 			fmt.Println(err)
// 		}
// 		noOfRows++
// 	}
// 	return noOfRows
// }
// func main() {
// 	var db Db
// 	r := db.ConnectToDb()
// 	selectQuery := `select * from mytodo;`
// 	res, err := r.Query(selectQuery)
// 	defer res.Close()

// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Println(handleRowsReturedByQuery(res))

// }
