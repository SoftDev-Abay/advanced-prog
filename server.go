package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

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

	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/message", MessageReciever)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
