package model

import (
	"fmt"
)

type todo struct {
	Id int
	Title string
	Content string
}

// type newTodo struct {
// 	Title string
// 	Content string
// }
func (t todo) createTodo(){
	
	fmt.Println("This is model")
}

func (t todo) deleteTodo(){
	// This will delete todo
}

func(t todo) updateTood(){
	// This will update todo
}

func (t todo) getAllTodo(){
	// this will get all todo
}