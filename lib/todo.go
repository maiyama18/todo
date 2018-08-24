package todolib

import "fmt"

type Todo struct {
	Title string
	Done bool
}

type Todos []Todo

func (todo Todo) PrintTodoLine(index int) {
	if todo.Done {
		fmt.Print("[x] ")
	} else {
		fmt.Print("[ ] ")
	}

	fmt.Printf("#%d %s\n", index, todo.Title)
}