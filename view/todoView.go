package view

import (
	"fmt"
	
)

func DisplayInterface() int {
	fmt.Println(("Choose no. based on operation you want to perform..."))
	fmt.Print("1. To add todo\n2. To get all todo\n3. To get a todo.\n4. To delete a todo\n5. To exit\n")
	var userChoise int
	fmt.Scanf("%d", &userChoise)
	return userChoise
}
