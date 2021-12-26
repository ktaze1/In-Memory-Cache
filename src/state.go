package main

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"os"
	"sync"
)

var lock sync.Mutex

var Marshal = func(v interface{}) (io.Reader, error) {
	b, err := json.MarshalIndent(v, "", "\t")
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(b), nil
}

func Save() error {
	lock.Lock()
	defer lock.Unlock()
	f, err := os.Create("./file.tmp")
	if err != nil {
		return err
	}
	defer f.Close()
	r, err := Marshal(store.m)
	if err != nil {
		return err
	}
	_, err = io.Copy(f, r)
	return err
}

func Load(path string, v interface{}) error {
	lock.Lock()
	defer lock.Unlock()
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()
	byteValue, _ := ioutil.ReadAll(f)

	return json.Unmarshal([]byte(byteValue), &store.m)
}
