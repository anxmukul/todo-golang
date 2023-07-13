package model

import (
	"fmt"
)

type Todo struct {
	Id int
	Title string
	Content string
}

// type newTodo struct {
// 	Title string
// 	Content string
// }
func (t Todo) CreateTodo(){
	fmt.Println("INside model createTod()");
}

func (t Todo) DleteTodo(){
	// This will delete todo
}

func(t Todo) UpdateTood(){
	// This will update todo
}

func (t Todo) getAllTodo(){
	// this will get all todo
}