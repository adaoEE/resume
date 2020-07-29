package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("."))))
	r.HandleFunc("/login", handler).Methods("POST")

	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

// change to create user, take in password, names, put it into database, maybe login
func handler(w http.ResponseWriter, r *http.Request) {
	u := &user{}
	if err := json.NewDecoder(r.Body).Decode(u); err != nil {
		panic(err)
	}
	fmt.Printf("%+v", u)
	rp := respone{"dinosaur 404"}
	json.NewEncoder(w).Encode(rp)
}

type respone struct {
	Message string
}

type user struct {
	Username string
	Password string
}
