package main

import (
	"bufio"
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

func getUserInput() string {
	titleScanner := bufio.NewScanner(os.Stdin)
	titleScanner.Scan()
	err := titleScanner.Err()
	if err != nil {
		panic(err)
	}
	return titleScanner.Text()
}

type todo struct {
	Id      int
	Title   string
	Content string
}

func (t todo) insertTodo() {
	t.Id = 0
	t.Title = getUserInput()
	t.Content = getUserInput()
	r := m.Todo{}
	r.CreateTodo(t)
}

func handleRequest(choice int) {
	switch choice {
	case 1:
		fmt.Printf("Your choose %d to create a todo\n", choice)
		t := todo{}
		t.insertTodo()
		v.DisplayInterface()
	case 2:
		fmt.Printf("Your choose %d to show all todo\n", choice)
		v.DisplayInterface()
	case 3:
		fmt.Printf("Your choose %d to show a todo\n", choice)
		v.DisplayInterface()
	case 4:
		fmt.Printf("Your choose %d to delete a todo\n", choice)
		v.DisplayInterface()
	case 5:
		fmt.Printf("You choose %d to exit.\nExiting.......\n", choice)
		os.Exit(0)
	default:
		fmt.Println("Invalid Choice.. choose again")
		v.DisplayInterface()
	}
}

func main() {
	fmt.Println("This is a Todo Application!")
	res := v.DisplayInterface()
	fmt.Println(res)
}
