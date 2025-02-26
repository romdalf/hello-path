package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func main() {
	// ensure log directory exists
	logDir := "http-hello-log"
	if err := os.MkdirAll(logDir, 0755); err != nil {
		log.Fatal("Failed to create log directory:", err)
	}

	// open log file with current date
	logFile := filepath.Join(logDir, time.Now().Format("2006-01-02")+".log")
	f, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal("Failed to open log file:", err)
	}
	defer f.Close()

	// create logger
	logger := log.New(f, "", log.LstdFlags)

	// handle HTTP requests
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello from URL path: %s\n", r.URL.Path)

		if r.URL.Path == "/" {
			fmt.Fprintf(w, "Try to add /netapp as a path.")
		}

		if r.URL.Path != "/favicon.ico" {
			// log to console
			fmt.Printf("User requested the URL path: %s\n", r.URL.Path)
			// log to file
			logger.Printf("Request from %s: %s %s\n", r.RemoteAddr, r.Method, r.URL.Path)
		}
	})

	fmt.Println("Hello from Path Service")
	fmt.Println("--> Server running on http://localhost:8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
