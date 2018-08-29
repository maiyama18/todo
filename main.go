package main

import (
	"os"
	"log"

	"github.com/urfave/cli"
	"path/filepath"
	"github.com/ymr-39/todo/db"
	"fmt"
	"strings"
	"errors"
	"strconv"
)

func main() {
	dir, err := getHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	dbPath := filepath.Join(dir, ".todos.db")
	db.Init(dbPath)

	app := cli.NewApp()
	app.Name = "todo"
	app.Version = "0.0.1"

	app.Commands = []cli.Command{
		{
			Name:  "add",
			Usage: "add a task",
			Action: func(c *cli.Context) error {
				todoTitle := strings.Join(c.Args(), " ")
				if todoTitle == "" {
					return errors.New("please enter non-empty todo")
				}

				return db.CreateTodo(todoTitle)
			},
		},
		{
			Name:  "list",
			Usage: "show a todo list",
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:  "undone, u",
					Usage: "show only undone todos",
				},
				cli.BoolFlag{
					Name:  "done, d",
					Usage: "show only done todos",
				},
			},
			Action: func(c *cli.Context) error {
				todos, err := db.AllTodos()
				if err != nil {
					return err
				}

				for _, todo := range todos {
					if c.Bool("undone") && todo.Done {
						continue
					} else if c.Bool("done") && !todo.Done {
						continue
					}

					fmt.Println(todo.TodoLine())
				}

				return nil
			},
		},
		{
			Name:  "done",
			Usage: "make a todo done",
			Action: func(c *cli.Context) error {
				idStr := c.Args().First()
				id, err := strconv.Atoi(idStr)
				if err != nil {
					return err
				}

				err = db.ToggleTodo(uint64(id), true)
				if err != nil {
					return err
				}

				return nil
			},
		},
		{
			Name:  "undone",
			Usage: "make a todo undone",
			Action: func(c *cli.Context) error {
				idstr := c.Args().First()
				id, err := strconv.Atoi(idstr)
				if err != nil {
					return err
				}

				err = db.ToggleTodo(uint64(id), false)
				if err != nil {
					return err
				}

				return nil
			},
		},
		{
			Name:  "remove",
			Usage: "remove a task",
			Action: func(c *cli.Context) error {
				for _, arg := range c.Args() {
					id, err := strconv.Atoi(arg)
					if err != nil {
						return err
					}

					db.DeleteTodo(uint64(id))
				}

				return nil
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
