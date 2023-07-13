package view

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type todos interface {
	getTodoTitleFromUser()
	getTodoFieldFromUser()
	showTodo()
	showAllTodo()
}
type Todo struct {
	Id      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

// func (t *Todo) ShowAllTodo() {

// }

func (t Todo) ShowTodo() {
	fmt.Println("Todo:")
	res, err := json.MarshalIndent(t, "", "")
	if err != nil {
		fmt.Println("Error while converting into JSON")
	}
	fmt.Println(string(res))

}
func (t Todo) GetTodoTitleFromUser() Todo {
	fmt.Printf("Enter title of Todo\n")
	titleScanner := bufio.NewScanner(os.Stdin)
	titleScanner.Scan()
	var err error
	err = titleScanner.Err()
	if err != nil {
		panic(err)
	}
	t.Title = titleScanner.Text()
	return t
}

func (t Todo) GetTodoFieldFromUser() Todo {
	fmt.Printf("Enter title of Todo\n")
	titleScanner := bufio.NewScanner(os.Stdin)
	titleScanner.Scan()
	var err error
	err = titleScanner.Err()
	if err != nil {
		panic(err)
	}
	t.Title = titleScanner.Text()
	fmt.Printf("Enter content of Todo\n")
	contentScanner := bufio.NewScanner(os.Stdin)
	contentScanner.Scan()
	err = contentScanner.Err()
	if err != nil {
		panic(err)
	}
	t.Content = contentScanner.Text()
	return t
}

func DisplayInterface() int {
	fmt.Println(("Choose no. based on operation you want to perform..."))
	fmt.Print("1. To add todo\n2. To get a todo\n3. To get all todo.\n4. To delete a todo\n5. To exit\n")
	var userChoise int
	fmt.Scanf("%d", &userChoise)
	return userChoise
}
