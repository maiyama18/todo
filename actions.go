package main

import (
	"github.com/urfave/cli"
	"fmt"
	"os"

	todo "github.com/m4iyama/todo/lib"
	"strconv"
)

func addAction(todos todo.Todos, jsonFilename string) func(c *cli.Context) error {
	return func(c *cli.Context) error {
		todoTitle := c.Args().First()
		if todoTitle == "" {
			fmt.Println("please enter non-empty todo")
			os.Exit(1)
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

			if todo.Done {
				fmt.Print("[x] ")
			} else {
				fmt.Print("[ ] ")
			}

			fmt.Printf("#%d %s\n", i, todo.Title)
		}

		return nil
	}
}

func doneAction(todos todo.Todos, jsonFilename string) func(c *cli.Context) error {
	return func(c *cli.Context) error {
		index, err := strconv.Atoi(c.Args().First())
		if err != nil || index < 0 || index > len(todos) {
			fmt.Println("please enter number of todo to done")
			os.Exit(1)
		}

		if todos[index].Done {
			fmt.Printf("todo #%d is already done", index)
			os.Exit(1)
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
			fmt.Println("please enter number of todo to undone")
			os.Exit(1)
		}

		if !todos[index].Done {
			fmt.Printf("todo #%d is still undone", index)
			os.Exit(1)
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
			fmt.Println("please enter number of todo to remove")
			return err
		}

		todos := append(todos[:index], todos[index+1:]...)

		if err := saveTodosToJsonfile(jsonFilename, todos); err != nil {
			return err
		}

		return nil
	}
}
