package main

import (
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	middleware "github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/mux"
)

func notAllowedHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Not Allowed", http.StatusMethodNotAllowed)
}

func setKeyHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]

	value, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = Set(key, string(value))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)

	log.Printf("PUT REQUEST key=%s value=%s\n", key, string(value))
}

func getKeyHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]

	value, err := Get(key)
	if errors.Is(err, ErrorNoSuchKey) {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte(value))

	log.Printf("GET REQUEST key=%s\n", key)
}

func keyValueDeleteHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]

	err := Delete(key)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Printf("DELETE key=%s\n", key)
}
func flushHandler(w http.ResponseWriter, r *http.Request) {
	err := Flush()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Printf("FLUSH")

}

func main() {

	if err := Load("./file.tmp", store.m); err != nil {
		log.Println(err)
	}
	t := schedule(Save, 20*time.Second)

	r := mux.NewRouter()

	r.Use(logger)

	r.HandleFunc("/v1/{key}", getKeyHandler).Methods("GET")

	r.HandleFunc("/v1/{key}", setKeyHandler).Methods("PUT")
	r.HandleFunc("/v1/{key}", keyValueDeleteHandler).Methods("DELETE")
	r.HandleFunc("/flush", flushHandler).Methods("DELETE")

	r.HandleFunc("/v1", notAllowedHandler)
	r.HandleFunc("/v1/{key}", notAllowedHandler)

	opts := middleware.SwaggerUIOpts{SpecURL: "./swagger.yaml"}
	sh := middleware.SwaggerUI(opts, nil)
	r.Handle("/docs", sh)
	log.Fatal(http.ListenAndServe(":8080", r))
	t.Stop()
}
