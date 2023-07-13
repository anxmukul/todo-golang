package main

import (
	"fmt"
	"os"

	d "github.com/anxmukul/todo-golang/database"
	m "github.com/anxmukul/todo-golang/model"
	v "github.com/anxmukul/todo-golang/view"
)

type todoInterface interface {
	insertTodo()
	showAllTodo()
	showTodo()
	deleteTodo()
}

type Todo struct {
	Id      int
	Title   string
	Content string
}

func (t Todo) createTodo() int {
	r := m.Todo{Id: t.Id, Title: t.Title, Content: t.Content}
	todoId := r.InsertTodo()
	return todoId
}

func (t Todo) getTodo() Todo {
	r := m.Todo{Title: t.Title}
	returedTodo := r.GetTodo()
	res := Todo{Id: returedTodo.Id, Title: returedTodo.Title, Content: returedTodo.Content}
	return res
}

func (t Todo) deleteTodo() int {
	r := m.Todo{Title: t.Title}
	todoId := r.DeleteTodo()
	return todoId
}

func (t Todo) showAllTodo() []Todo {
	r := m.Todo{}
	res := r.GetAllTodo()
	var todoArray = make([]Todo, 0)
	for _, item := range res {
		newTodo := Todo{Id: item.Id, Title: item.Title, Content: item.Content}
		todoArray = append(todoArray, newTodo)
	}
	return todoArray
}

func showall(t []Todo) {
	var todoArray = make([]v.Todo, 0)
	for _, item := range t {
		newTodo := v.Todo{Id: item.Id, Title: item.Title, Content: item.Content}
		todoArray = append(todoArray, newTodo)
		newTodo.ShowTodo()
	}
}

// func convertIntoViewTodoArray(t []Todo) *[]v.Todo {
// 	var todoArray = make([]v.Todo, 0)
// 	for _, item := range t {
// 		newTodo := v.Todo{Id: item.Id, Title: item.Title, Content: item.Content}
// 		todoArray = append(todoArray, newTodo)
// 	}
// 	// obj.ShowAllTodo(&todoArray)
// 	return &todoArray
// }

func handleRequest(choice int) {
	switch choice {
	case 1:
		// fmt.Printf("Your choose %d to create a Todo\n", choice)
		myTodo := v.Todo{}
		res := myTodo.GetTodoFieldFromUser()
		t := Todo{Id: res.Id, Title: res.Title, Content: res.Content}
		id := t.createTodo()
		fmt.Printf("Todo created with Id %d\n", id)
		handleRequest(v.DisplayInterface())
	case 2:
		// fmt.Printf("Your choose %d to show Todo\n", choice)
		myTodo := v.Todo{}
		input := myTodo.GetTodoTitleFromUser()
		t := Todo{Title: input.Title}
		returedTodo := t.getTodo()
		res := v.Todo{Id: returedTodo.Id, Title: returedTodo.Title, Content: returedTodo.Content}
		res.ShowTodo()
		handleRequest(v.DisplayInterface())
	case 3:
		// fmt.Printf("Your choose %d to show all Todo\n", choice)
		myTodo := Todo{}
		returedTodo := myTodo.showAllTodo()
		// res := convertIntoViewTodoArray(returedTodo)
		showall(returedTodo)
		handleRequest(v.DisplayInterface())
	case 4:
		// fmt.Printf("Your choose %d to delete a Todo\n", choice)
		myTodo := v.Todo{}
		input := myTodo.GetTodoTitleFromUser()
		t := Todo{Title: input.Title}
		res := t.deleteTodo()
		fmt.Printf("Todo with id %d is deleted\n", res)
		handleRequest(v.DisplayInterface())
	case 5:
		fmt.Printf("You choose %d to exit.\n", choice)
		os.Exit(0)
	default:
		fmt.Println("Invalid Choice.. choose again")
		handleRequest(v.DisplayInterface())
	}
}

func main() {
	fmt.Println("This is a Todo Application!")
	var database d.Db
	db := database.ConnectToDb()
	err := db.Ping()
	if err != nil {
		fmt.Println("Database connection failed")
		os.Exit(1)
	}
	res := v.DisplayInterface()
	handleRequest(res)
}
