package db

import (
	"errors"
	"sync"
)

type (
	KeyValueStore interface {
		Get(key string) (value string, err error)
		Set(key, value string) error
	}

	kvs struct {
		mux  sync.Mutex
		data map[string]string
	}
)

var ErrKeyNotFound = errors.New("key not found")

func NewKeyValueStore() KeyValueStore {
	return &kvs{data: make(map[string]string)}
}

func (d *kvs) Get(key string) (value string, err error) {
	value, ok := d.data[key]
	if !ok {
		return "", ErrKeyNotFound
	}

	return value, nil
}

func (d *kvs) Set(key string, value string) error {
	d.mux.Lock()
	defer d.mux.Unlock()

	d.data[key] = value
	return nil
}
