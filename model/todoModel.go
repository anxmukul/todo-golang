package model

import (
	"fmt"

	db "github.com/anxmukul/todo-golang/database"
)

type Todo struct {
	Id      int
	Title   string
	Content string
}

// type newTodo struct {
// 	Title string
// 	Content string
// }

func (t Todo) InsertTodo(){
	fmt.Println("Inside model createTodo()")
	fmt.Printf("{Title: %s\tContent: %s}\n", t.Title, t.Content)
	var d db.Db
	db := d.ConnectToDb()
	insertQuery := `insert into mytodo(title, content) values($1, $2) returning id;`
	var id int
	err := db.QueryRow(insertQuery, t.Title, t.Content).Scan(&id)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Createdted to with id:", id)
}

func (t Todo) DeleteTodo() {
	// This will delete todo
	fmt.Println("Inside model DeleteTodo()")

}

func (t Todo) GetTodo() {
	// This will update todo
	fmt.Println("Inside model GetTodo()")

}

func (t Todo) GetAllTodo() {
	// this will get all todo
	fmt.Println("Inside model GetAllTodo()")

}
