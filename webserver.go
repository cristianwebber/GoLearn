package main

import (
	"net/http"
	"text/template";
	"encoding/json";
)

type Task struct{
	Name string
	Done bool
}

func main() {
	http.HandleFunc("/", Hello)
	http.HandleFunc("/tasks", TaskHandler)
	http.HandleFunc("/api", ApiHandler)
	http.ListenAndServe(":8888", nil)
}

func Hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

func TaskHandler(w http.ResponseWriter, r *http.Request) {

	var task Task = Task{
		Name: "Learn Go",
		Done: false,
	}

	t := template.Must(template.ParseFiles("task.html"))
	t.Execute(w, task)
}

func ApiHandler(w http.ResponseWriter, r *http.Request) {

	var task Task = Task{
		Name: "Learn Go",
		Done: false,
	}

	j, _ := json.Marshal(task)
	w.Write(j)
}