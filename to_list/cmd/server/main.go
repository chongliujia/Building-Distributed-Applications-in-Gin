package main

import (
    "log"
    "net/http"
    "to_list/internal/api"
)

func main() {
    http.HandleFunc("/todos", api.ListTodosHandler)
    http.HandleFunc("/add", api.AddTodoHandler)

    port := ":8080"
    log.Printf("Server starting on port %s\n", port)

    if err := http.ListenAndServe(port, nil); err != nil {
        log.Fatalf("Failed to start server: %v", err)
    }
}
