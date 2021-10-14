package db

import "encoding/json"

type User struct {
	Name           string  `json:"name"`
	HashedPassword string  `json:"hashed_password"`
	ToDos          []*ToDo `json:"todos"`
}

var keyPrefixUser = "/users/"

func (d *User) Save(store KeyValueStore) error {
	data, err := json.Marshal(d)
	if err != nil {
		return err
	}

	if err := store.Set(keyPrefixUser+d.Name, string(data)); err != nil {
		return err
	}
	return nil
}

func (d *User) Get(store KeyValueStore) error {
	data, err := store.Get(keyPrefixUser + d.Name)
	if err != nil {
		return err
	}

	if err := json.Unmarshal([]byte(data), d); err != nil {
		return err
	}
	return nil
}
