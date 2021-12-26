package main

import (
	"errors"
	"fmt"
	"sync"
)

var ErrorNoSuchKey = errors.New("given key not found")

var store = struct {
	sync.RWMutex
	m map[string]string
}{m: make(map[string]string)}

func Delete(key string) error {
	store.Lock()
	delete(store.m, key)
	store.Unlock()

	return nil
}

func Flush() error {
	store.Lock()
	fmt.Println("Flushed")
	store.m = make(map[string]string)
	store.Unlock()

	return nil
}
func Get(key string) (string, error) {
	store.RLock()
	value, ok := store.m[key]
	store.RUnlock()

	if !ok {
		return "", ErrorNoSuchKey
	}

	return value, nil
}

func Set(key string, value string) error {
	store.Lock()
	store.m[key] = value
	store.Unlock()

	return nil
}
