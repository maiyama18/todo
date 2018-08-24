package todolib

import "fmt"

type Todo struct {
	Title string
	Done bool
}

type Todos []Todo

func (todo Todo) TodoLine(index int) string {
	var check string
	if todo.Done {
		check = "[x]"
	} else {
		check = "[ ]"
	}

	return fmt.Sprintf("%s #%d %s", check, index, todo.Title)
}