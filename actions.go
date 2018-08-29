package main

import (
	"github.com/urfave/cli"
	"fmt"
	todo "github.com/ymr-39/todo/lib"
	"strconv"
	"errors"
)

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

		fmt.Printf("todo: '%s' is done\n", todos[index].Title)

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

		fmt.Printf("todo: '%s' is now undone\n", todos[index].Title)

		return nil
	}
}

func removeAction(todos todo.Todos, jsonFilename string) func(c *cli.Context) error {
	return func(c *cli.Context) error {
		index, err := strconv.Atoi(c.Args().First())
		if err != nil || index < 0 || index > len(todos) {
			return errors.New("please enter valid index of todo")
		}

		todoTitle := todos[index].Title

		todos := append(todos[:index], todos[index+1:]...)

		if err := saveTodosToJsonfile(jsonFilename, todos); err != nil {
			return err
		}

		fmt.Printf("todo: '%s' was removed\n", todoTitle)

		return nil
	}
}
