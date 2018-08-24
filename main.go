package main

import (
	"fmt"
	"os"
	"log"

	"github.com/urfave/cli"
	"encoding/json"
	"io/ioutil"
	todo "github.com/m4iyama/todo/lib"
	"os/user"
	"path/filepath"
	"strconv"
)

func main() {
	dir, err := getHomeDir()
	if err != nil {
		panic(err)
	}

	jsonFilename := filepath.Join(dir, ".todos.json")
	if _, err := os.Stat(jsonFilename); os.IsNotExist(err) {
		if err := createEmptyJsonfile(jsonFilename); err != nil {
			panic(err)
		}
	}

	f, err := os.Open(jsonFilename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	bytes, err := ioutil.ReadAll(f)
	var todos todo.Todos
	if err := json.Unmarshal(bytes, &todos); err != nil {
		panic(err)
	}

	app := cli.NewApp()
	app.Name = "todo"
	app.Version = "0.0.1"

	app.Commands = []cli.Command{
		{
			Name:  "add",
			Usage: "add a task",
			Action: func(c *cli.Context) error {
				todoTitle := c.Args().First()
				if todoTitle == "" {
					fmt.Println("please enter non-empty todo")
					os.Exit(1)
				}

				todos := append(todos, todo.Todo{Title: todoTitle, Done: false})
				fmt.Println(todos)

				if err := saveTodosToJsonfile(jsonFilename, todos); err != nil {
					return err
				}

				return nil
			},
		},
		{
			Name:  "list",
			Usage: "show a todo list",
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name: "undone, u",
					Usage: "show only undone todos",
				},
			},
			Action: func(c *cli.Context) error {
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
			},
		},
		{
			Name:  "done",
			Usage: "make a todo done",
			Action: func(c *cli.Context) error {
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
			},
		},
		{
			Name:  "undone",
			Usage: "make a todo undone",
			Action: func(c *cli.Context) error {
				index, err := strconv.Atoi(c.Args().First())
				if err != nil || index < 0 || index > len(todos) {
					fmt.Println("please enter number of todo to undone")
					os.Exit(1)
				}

				if todos[index].Done {
					fmt.Printf("todo #%d is still undone", index)
					os.Exit(1)
				}

				todos[index].Done = false

				if err := saveTodosToJsonfile(jsonFilename, todos); err != nil {
					return err
				}

				return nil
			},
		},
		{
			Name:  "remove",
			Usage: "remove a task",
			Action: func(c *cli.Context) error {
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
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func getHomeDir() (string, error) {
	usr, err := user.Current()
	if err != nil {
		return "", err
	}

	return usr.HomeDir, nil
}

func createEmptyJsonfile(jsonFilename string) error {
	f, err := os.Create(jsonFilename)
	if err != nil {
		return err
	}
	defer f.Close()

	f.Write([]byte("[]"))

	return nil
}

func saveTodosToJsonfile(jsonFilename string, todos todo.Todos) error {
	bytes, err := json.Marshal(todos)
	if err != nil {
		return err
	}

	if err := os.Remove(jsonFilename); err != nil {
		return err
	}

	ff, err := os.Create(jsonFilename)
	if err != nil {
		return err
	}
	defer ff.Close()

	if _, err = ff.Write(bytes); err != nil {
		return err
	}

	return nil
}
