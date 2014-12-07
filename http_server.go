package main

import (
	"io"
	"net/http"
	"log"
)

func HelloServer(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello, GDG!\n")
}

func main() {
	http.HandleFunc("/", HelloServer)
	log.Println("Listening on 0.0.0.0:8081")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
