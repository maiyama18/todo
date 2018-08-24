package main

import (
	"os/user"
	"os"
	"encoding/json"

	todo "github.com/ymr-39/todo/lib"
)

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
