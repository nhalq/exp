package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		return "2871"
	}

	return port
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(5 * time.Second)
		fmt.Fprintf(w, "%s ok", r.URL.Path)
	})

	port := ":" + getPort()
	fmt.Println("Starting server on " + port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Println(err)
	}
}
