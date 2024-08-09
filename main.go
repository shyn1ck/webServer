package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Task struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	IsDone      bool   `json:"is_done"`
	IsDeleted   bool   `json:"is_deleted"`
	Priority    int    `json:"priority"`
}

type DefaultResponse struct {
	Message string `json:"message"`
}

var tasks []Task

func main() {

	http.HandleFunc("/get-all-tasks", getAllTasks)
	http.HandleFunc("/add-task", AddTask)
	fmt.Println("Server is listening on port 8585...")
	err := http.ListenAndServe(":8585", nil)
	if err != nil {
		panic(err)
	}
}

func getAllTasks(w http.ResponseWriter, r *http.Request) {

	t1 := Task{
		ID:          1,
		Title:       "Выучить GO",
		Description: "Занятся программирование",
		IsDone:      false,
		IsDeleted:   false,
		Priority:    1,
	}
	tasks = append(tasks, t1)
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(tasks)
	if err != nil {
		http.Error(w, "error while writing", http.StatusInternalServerError)
		return
	}
}

func AddTask(w http.ResponseWriter, r *http.Request) {
	var t Task
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	tasks = append(tasks, t)
	var response DefaultResponse
	response.Message = "Successfully added task"
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "error while writing", http.StatusInternalServerError)
		return
	}
}
