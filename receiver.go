package main

import (
	"net/http"
	"log"
)
func main() {
	http.HandleFunc("/", receiveImage)
	log.Print("launch the server")
	if err := http.ListenAndServe(":4321", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func receiveImage(w http.ResponseWriter, r *http.Request) {
	log.Print("retrieve the image here")
}