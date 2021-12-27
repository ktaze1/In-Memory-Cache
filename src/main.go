//  Api:
//   version: 0.0.1
//   title: Key-Value Store Api
//  Schemes: http, https
//  Host: localhost:8080
//  BasePath: /
//  Produces:
//    - application/json
//
// securityDefinitions:
//  apiKey:
//    type: apiKey
//    in: header
//    name: authorization
// swagger:meta
package main

import (
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/rs/cors"

	middleware "github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/mux"
)

func notAllowedHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Not Allowed", http.StatusMethodNotAllowed)
}

// swagger:route PUT /v1/ value
// Add record of given key
// consumes:
//  - text/plain
// responses:
//	201: commonResponse
//  404: commonResponse
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

// swagger:route GET /v1/ key
// Get record of given key
//
// responses:
//	200: commonResponse
//  404: commonResponse
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

// swagger:route DELETE /v1/ key
// Delete record of given key
//
// responses:
//	200: commonResponse
//  404: commonResponse
//  500: commonResponse
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

// swagger:route DELETE /flush key
// Delete record of given key
//
// responses:
//	200: commonResponse
//  500: commonResponse
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
	t := schedule(Save, 10*time.Minute)

	r := mux.NewRouter()

	r.Use(logger)

	r.HandleFunc("/v1/{key}", getKeyHandler).Methods("GET")
	r.HandleFunc("/v1/{key}", setKeyHandler).Methods("PUT")
	r.HandleFunc("/v1/{key}", keyValueDeleteHandler).Methods("DELETE")
	r.HandleFunc("/flush", flushHandler).Methods("DELETE")

	r.HandleFunc("/v1", notAllowedHandler)
	r.HandleFunc("/v1/{key}", notAllowedHandler)

	// Swagger
	opts := middleware.SwaggerUIOpts{SpecURL: "./swagger.yaml"}
	sh := middleware.SwaggerUI(opts, nil)
	r.Handle("/docs", sh).Methods("GET", "OPTIONS")

	r.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))
	corshandler := cors.AllowAll().Handler(r)

	log.Fatal(http.ListenAndServe(":8080", corshandler))
	
	t.Stop()
}
