package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	db "renting/database"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "Hello!")
}

func formHandler(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/form" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	if r.Method != "POST" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "POST request successful")
	name := r.FormValue("name")
	address := r.FormValue("address")

	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Address = %s\n", address)
}

type message struct {
	Message string `json:"message"`
}

type messageResponce struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func MessageReciever(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/message" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	if r.Method != "POST" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	decoder := json.NewDecoder(r.Body)
	var m message
	err := decoder.Decode(&m)

	mR := messageResponce{"", ""}

	if err != nil || m.Message == "" {
		mR.Status = "400"
		mR.Message = "Invalid JSON message"
		log.Println(err)
	} else {
		mR.Status = "success"
		mR.Message = "Data successfully received"
		log.Println(m.Message)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(mR)

}

func main() {

	fileServer := http.FileServer(http.Dir("./static")) // New code
	http.Handle("/", fileServer)                        // New code
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/message", MessageReciever)

	fmt.Printf("Starting server at port 8080\n")
	var arr []string

	arr, err := db.GetBuildings()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(arr)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}
