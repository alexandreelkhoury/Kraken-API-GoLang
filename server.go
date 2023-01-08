package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
)

func getTimee(w http.ResponseWriter, r *http.Request) {
	// Connect to the database
	db, err := sql.Open("postgres", "user=postgres password=mypassword dbname=krakendb sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Prints the time from the database on localhost:8080/time
	var rfc1123 string
	err = db.QueryRow("SELECT rfc1123 FROM time ORDER BY unixtime DESC LIMIT 1").Scan(&rfc1123)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		log.Fatal(err)
		return
	}
	w.Write([]byte(rfc1123))
}

func getStatuss(w http.ResponseWriter, r *http.Request) {
	// Connect to the database
	db, err := sql.Open("postgres", "user=postgres password=mypassword dbname=krakendb sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Prints the status from the database on localhost:8080/status
	var status string
	err = db.QueryRow("SELECT status FROM status ORDER BY timestamp DESC LIMIT 1").Scan(&status)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		log.Fatal(err)
		return
	}
	w.Write([]byte(status))
}

func startServer() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Kraken API")
	})
	http.HandleFunc("/time", getTimee)
	http.HandleFunc("/status", getStatuss)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
