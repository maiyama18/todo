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
		//{
		//	Name:  "done",
		//	Usage: "make a todo done",
		//	Action: doneAction(todos, dbPath),
		//},
		//{
		//	Name:  "undone",
		//	Usage: "make a todo undone",
		//	Action: undoneAction(todos, dbPath),
		//},
		//{
		//	Name:  "remove",
		//	Usage: "remove a task",
		//	Action: removeAction(todos, dbPath),
		//},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
