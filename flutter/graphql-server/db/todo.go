package db

import (
	"github.com/google/uuid"
)

type ToDo struct {
	TodoID   string `json:"todo_id"`
	UserName string `json:"user_name"`
	Text     string `json:"text"`
	Done     bool   `json:"done"`
}

func (d *ToDo) Save(store KeyValueStore) error {
	if d.TodoID == "" {
		d.TodoID = uuid.NewString()
	}

	user := &User{Name: d.UserName}
	if err := user.Get(store); err != nil {
		return err
	}

	user.ToDos = append(user.ToDos, d)
	if err := user.Save(store); err != nil {
		return err
	}

	return nil
}
