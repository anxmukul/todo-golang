package main

import (
	"fmt"
	"os"

	m "github.com/anxmukul/todo-golang/model"
	v "github.com/anxmukul/todo-golang/view"
)

type todoInterface interface {
	insertTodo()
	showAllTodo()
	showTodo()
	deleteTodo()
}

type todo struct {
	Id      int
	Title   string
	Content string
}

func (t todo) createTodo() {
	r := m.Todo{Id: t.Id, Title: t.Title, Content: t.Content}
	r.InsertTodo()
}

func (t todo) getTodo() todo {
	// t.Title = getUserInput("title");
	r := m.Todo{Title: t.Title}
	returedTodo := r.GetTodo()
	res := todo{Id: returedTodo.Id, Title: returedTodo.Title, Content: returedTodo.Content}
	return res
}


func handleRequest(choice int) {
	switch choice {
	case 1:
		fmt.Printf("Your choose %d to create a todo\n", choice)
		myTodo := v.Todo{}
		res := myTodo.GetTodoFieldFromUser()
		t := todo{Id: res.Id, Title: res.Title, Content: res.Content}
		t.createTodo()
		handleRequest(v.DisplayInterface())
	case 2:
		fmt.Printf("Your choose %d to show todo\n", choice)
		myTodo := v.Todo{}
		input := myTodo.GetTodoTitleFromUser()
		t := todo{Title: input.Title}
		returedTodo := t.getTodo()
		res := v.Todo{Id: returedTodo.Id, Title: returedTodo.Title, Content: returedTodo.Content}
		res.ShowTodo()
		handleRequest(v.DisplayInterface())
	case 3:
		fmt.Printf("Your choose %d to show a todo\n", choice)
		handleRequest(v.DisplayInterface())
	case 4:
		fmt.Printf("Your choose %d to delete a todo\n", choice)
		r := m.Todo{}
		r.DeleteTodo()
		handleRequest(v.DisplayInterface())
	case 5:
		fmt.Printf("You choose %d to exit.\nExiting.......\n", choice)
		os.Exit(0)
	default:
		fmt.Println("Invalid Choice.. choose again")
		handleRequest(v.DisplayInterface())
	}
}

func main() {
	fmt.Println("This is a Todo Application!")
	// fmt.Println("connecting to Db...")
	// var database d.Db
	// db := database.ConnectToDb()
	// fmt.Println(db)
	res := v.DisplayInterface()
	handleRequest(res)
}
