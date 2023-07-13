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

func (t Todo) InsertTodo() int {
	var d db.Db
	db := d.ConnectToDb()
	insertQuery := `insert into mytodo(title, content) values($1, $2) returning id;`
	var id int
	err := db.QueryRow(insertQuery, t.Title, t.Content).Scan(&id)
	if err != nil {
		fmt.Println(err)
	}
	return id
}

func (t Todo) DeleteTodo() int {
	// This will delete todo
	fmt.Println("Inside model DeleteTodo()")
	var id int
	var d db.Db
	db := d.ConnectToDb()
	deleteQuery := `delete from mytodo where title = $1 returning id`
	err := db.QueryRow(deleteQuery, t.Title).Scan(&id)
	if err != nil {
		fmt.Println(err)
	}
	return id

}

func (t Todo) GetTodo() Todo {
	var d db.Db
	db := d.ConnectToDb()
	selectQuery := `select * from mytodo where title = $1;`
	var res Todo
	err := db.QueryRow(selectQuery, t.Title).Scan(&res.Id, &res.Title, &res.Content)
	if err != nil {
		fmt.Println(err)
	}
	return res
}

func (t Todo) GetAllTodo() []Todo {
	fmt.Println("Inside model GetAllTodo()")
	selectQuery := `select * from mytodo;`
	var d db.Db
	db := d.ConnectToDb()
	rows, err := db.Query(selectQuery)
	defer rows.Close()
	if err != nil {
		fmt.Println(err)
	}
	var todoArray = make([]Todo, 0)
	for rows.Next() {
		var newTodo Todo
		err := rows.Scan(&newTodo.Id, &newTodo.Title, &newTodo.Content)
		if err != nil {
			fmt.Println(err)
		}
		todoArray = append(todoArray, newTodo)
	}
	return todoArray
}
