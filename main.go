package main

import (
	"fmt"
	"os"
	v "github.com/anxmukul/todo-golang/view"
)



func handleRequest(choice int) {
	switch choice {
	case 1:
		fmt.Printf("Your choose %d \n", choice)
		v.DisplayInterface()
	case 2:
		fmt.Printf("Your choose %d \n", choice)
		v.DisplayInterface()
	case 3:
		fmt.Printf("Your choose %d \n", choice)
		v.DisplayInterface()
	case 4:
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
