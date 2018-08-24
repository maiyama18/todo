package main

import (
	"github.com/urfave/cli"
	"fmt"
	todo "github.com/m4iyama/todo/lib"
	"strconv"
	"errors"
)

func addAction(todos todo.Todos, jsonFilename string) func(c *cli.Context) error {
	return func(c *cli.Context) error {
		todoTitle := c.Args().First()
		if todoTitle == "" {
			return errors.New("please enter non-empty todo")
		}

		todos := append(todos, todo.Todo{Title: todoTitle, Done: false})

		if err := saveTodosToJsonfile(jsonFilename, todos); err != nil {
			return err
		}

		return nil
	}
}

func listAction(todos todo.Todos, _ string) func(c *cli.Context) error {
	return func(c *cli.Context) error {
		for i, todo := range todos {
			if c.Bool("undone") && todo.Done {
				continue
			}

			todo.PrintTodoLine(i)
		}

		return nil
	}
}

func doneAction(todos todo.Todos, jsonFilename string) func(c *cli.Context) error {
	return func(c *cli.Context) error {
		index, err := strconv.Atoi(c.Args().First())
		if err != nil || index < 0 || index > len(todos) {
			return errors.New("please enter valid index of todo")
		}

		if todos[index].Done {
			return errors.New(fmt.Sprintf("todo #%d is already done", index))
		}

		todos[index].Done = true

		if err := saveTodosToJsonfile(jsonFilename, todos); err != nil {
			return err
		}

		return nil
	}
}

func undoneAction(todos todo.Todos, jsonFilename string) func(c *cli.Context) error {
	return func(c *cli.Context) error {
		index, err := strconv.Atoi(c.Args().First())
		if err != nil || index < 0 || index > len(todos) {
			return errors.New("please enter valid index of todo")
		}

		if !todos[index].Done {
			return errors.New(fmt.Sprintf("todo #%d is still undone", index))
		}

		todos[index].Done = false

		if err := saveTodosToJsonfile(jsonFilename, todos); err != nil {
			return err
		}

		return nil
	}
}

func removeAction(todos todo.Todos, jsonFilename string) func(c *cli.Context) error {
	return func(c *cli.Context) error {
		index, err := strconv.Atoi(c.Args().First())
		if err != nil || index < 0 || index > len(todos) {
			return errors.New("please enter valid index of todo")
		}

		todos := append(todos[:index], todos[index+1:]...)

		if err := saveTodosToJsonfile(jsonFilename, todos); err != nil {
			return err
		}

		return nil
	}
}
