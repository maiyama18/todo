package main

import (
	"os"
	"log"

	"github.com/urfave/cli"
	"encoding/json"
	"io/ioutil"
	todo "github.com/ymr-39/todo/lib"
	"path/filepath"
)

func main() {
	dir, err := getHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	jsonFilename := filepath.Join(dir, ".todos.json")
	if _, err := os.Stat(jsonFilename); os.IsNotExist(err) {
		if err := createEmptyJsonfile(jsonFilename); err != nil {
			log.Fatal(err)
		}
	}

	f, err := os.Open(jsonFilename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	bytes, err := ioutil.ReadAll(f)
	var todos todo.Todos
	if err := json.Unmarshal(bytes, &todos); err != nil {
		log.Fatal(err)
	}

	app := cli.NewApp()
	app.Name = "todo"
	app.Version = "0.0.1"

	app.Commands = []cli.Command{
		{
			Name:  "add",
			Usage: "add a task",
			Action: addAction(todos, jsonFilename),
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
			Action: listAction(todos, jsonFilename),
		},
		{
			Name:  "done",
			Usage: "make a todo done",
			Action: doneAction(todos, jsonFilename),
		},
		{
			Name:  "undone",
			Usage: "make a todo undone",
			Action: undoneAction(todos, jsonFilename),
		},
		{
			Name:  "remove",
			Usage: "remove a task",
			Action: removeAction(todos, jsonFilename),
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

