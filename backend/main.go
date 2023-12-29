package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", getTime)
	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}

func getTime(w http.ResponseWriter, r *http.Request) {
	currentTime := time.Now().Format(time.RFC822)
	fmt.Fprintf(w, "%s", currentTime)
}
