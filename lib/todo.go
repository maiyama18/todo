package todolib

import "fmt"

type Todo struct {
	Id uint64
	Title string
	Done bool
}

type Todos []Todo

func (todo Todo) TodoLine() string {
	var check string
	if todo.Done {
		check = "[x]"
	} else {
		check = "[ ]"
	}

	return fmt.Sprintf("%s #%d %s", check, todo.Id, todo.Title)
}