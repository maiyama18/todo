package db

import (
	"github.com/boltdb/bolt"
	"time"
	todo "github.com/ymr-39/todo/lib"
	"encoding/binary"
	"encoding/json"
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

func btoi(b []byte) uint64 {
	return binary.BigEndian.Uint64(b)
}

func itob(i uint64) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, i)

	return b
}