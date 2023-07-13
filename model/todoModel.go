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
	fmt.Println("Inside model createTodo()");
	fmt.Printf("{Title: %s\tContent: %s}\n", t.Title, t.Content)
}

func (t Todo) DeleteTodo(){
	// This will delete todo
	fmt.Println("Inside model DeleteTodo()");

}

func(t Todo) GetTodo(){
	// This will update todo
	fmt.Println("Inside model GetTodo()");

}

func (t Todo) GetAllTodo(){
	// this will get all todo
	fmt.Println("Inside model GetAllTodo()");

}