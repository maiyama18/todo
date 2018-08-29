package db

import (
	"github.com/boltdb/bolt"
	"time"
	todo "github.com/ymr-39/todo/lib"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
)

var bucketName = []byte("todos")
var db *bolt.DB

func Init(dbPath string) error {
	var err error
	db, err = bolt.Open(dbPath, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return err
	}

	return db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(bucketName)
		return err
	})
}

func CreateTodo(title string) error {
	return db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucketName)
		id, _ := b.NextSequence()
		todo := todo.Todo{
			Id: id,
			Title: title,
			Done: false,
		}

		todoBytes, err := json.Marshal(todo)
		if err != nil {
			return err
		}

		return b.Put(itob(id), todoBytes)
	})
}

func AllTodos() (todo.Todos, error) {
	var todos todo.Todos
	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucketName)
		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			var todo todo.Todo
			json.Unmarshal(v, &todo)
			todos = append(todos, todo)
		}

		return nil
	})

	return todos, nil
}

func DeleteTodo(id uint64) error {
	return db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucketName)
		return b.Delete(itob(id))
	})
}

func ToggleTodo(id uint64, done bool) error {
	return db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucketName)
		key := itob(id)

		bytes := b.Get(key)
		todo, err := btot(bytes)
		if err != nil {
			return err
		} else if todo.Done == done {
			var state string
			if done {
				state = "already done"
			} else {
				state = "still undone"
			}
			return errors.New(fmt.Sprintf("todo #%d is %s", id, state))
		}

		err = b.Delete(key)
		if err != nil {
			return err
		}

		todo.Done = done
		bytes, err = ttob(todo)
		if err != nil {
			return err
		}

		return b.Put(key, bytes)
	})
}

func btoi(b []byte) uint64 {
	return binary.BigEndian.Uint64(b)
}

func itob(i uint64) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, i)

	return b
}

func btot(b []byte) (todo.Todo, error) {
	var t todo.Todo
	err := json.Unmarshal(b, &t)
	if err != nil {
		return todo.EmptyTodo(), err
	}

	return t, nil
}

func ttob(t todo.Todo) ([]byte, error) {
	var b []byte
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}

	return b, nil
}
