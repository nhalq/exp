package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sync"
	"time"
)

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		return "2871"
	}

	return port
}

func asyncRequest(wg *sync.WaitGroup, path string) {
	defer wg.Done()
	request, err := http.NewRequest("GET", "http://0.0.0.0:"+getPort()+"/"+path, nil)
	if err != nil {
		log.Fatalf("request: %s - error: %s", path, err)
	}

	client := http.Client{Timeout: 10 * time.Second}

	log.Printf("sending request: %s", path)
	response, _ := client.Do(request)
	defer response.Body.Close()

	content, _ := io.ReadAll(response.Body)
	log.Printf(string(content))
}

func main() {
	log.Printf("watch -n 1 ls -lhav /proc/%d/task", os.Getpid())
	log.Println("Press any key to continue...")
	fmt.Scanln()

	wg := sync.WaitGroup{}
	for i := 0; i < 1024; i++ {
		wg.Add(1)
		go asyncRequest(&wg, fmt.Sprintf("pending/%04d", i))
	}

	wg.Wait()
}
