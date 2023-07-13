package main

import(
	"fmt"
	"database/sql"
	_ "github.com/lib/pq"

)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "blah"
	dbname   = "todo"
)

func connectToDb() *sql.DB {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlconn)

	if err != nil {
		fmt.Println(err)
	}
	err = db.Ping()
	if err != nil {
		fmt.Println("Error occurd", err)
	}
	return db
}