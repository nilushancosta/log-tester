package main

import (
	"fmt"
	"net/http"
	"nilushancosta/log-tester/internal/pkg/handlers"
)

func main() {
	fmt.Println("Starting server on port 8080...")
	http.HandleFunc("/generate-logs-of-size", handlers.GenerateLogsOfSize)
	http.ListenAndServe(":8080", nil)
}
