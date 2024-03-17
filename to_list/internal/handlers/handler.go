package api

import (
    "encoding/json"
    "net/http"
    "sync"
    "to_list/internal/model"
)

var (
    todos  []model.Todo
    mutex  sync.Mutex
)


func TodosHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
        return 
    }

    var newTodo model.Todo
    err := json.NewDecoder(r.Body).Decode(&newTodo)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return 
    }

    mutex.Lock()
    newTodo.ID = len(todos) + 1
    todos = append(todos, newTodo)
    mutex.Unlock()

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(newTodo)
}

func ListTodosHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodGet {
        http.Error(w, "Only GET method is allowed", http.StatusMethodNotAllowed)
        return 
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(todos)
}
