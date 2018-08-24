package todolib

import (
	"testing"
	"fmt"
)

func TestTodoLine(t *testing.T) {
	todo := Todo{Title: "test", Done: false}
	index := 1

	expected := "[ ] #1 test"
	if actual := todo.TodoLine(index); actual != expected {
		t.Fatal(fmt.Sprintf("expected: %s, actual: %s", expected, actual))
	}
}
