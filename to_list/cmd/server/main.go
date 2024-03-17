package main

import (
    "net/http"
    "path/filepath"
    api "to_list/internal/handlers"
    "log"
)

func main() {
    static_path := filepath.Join("internal", "views")
    static_handler := http.FileServer(http.Dir(static_path))

    http.Handle("/", static_handler)

    http.HandleFunc("/api/todos", api.TodosHandler)

    log.Println("Server starting on port 8080...")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal(err)
    } 
}
