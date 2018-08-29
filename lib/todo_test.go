package todolib

import (
	"testing"
	"fmt"
)

func TestTodoLine(t *testing.T) {
	todo := Todo{Id: 1, Title: "test", Done: false}

	expected := "[ ] #1 test"
	if actual := todo.TodoLine(); actual != expected {
		t.Fatal(fmt.Sprintf("expected: %s, actual: %s", expected, actual))
	}
}
