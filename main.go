// var store = make(map[string]string)

// func Put(key string, value string) error {
// 	store[key] = value

// 	return nil
// }

// var ErrorNoSuchKey = errors.New("no such key")

// func Get(key string) (string, error) {
// 	value, ok := store[key]

// 	if !ok {
// 		return "", ErrorNoSuchKey
// 	}

// 	return value, nil
// }

// func Delete(key string) error {
// 	delete(store, key)

// 	return nil
// }

package main

import (
	"log"
	"net/http"
)

func HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

func main() {
	http.HandleFunc("/", HelloWorldHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
