package main

import (
	"html/template"
	"log"
	"net/http"
)

type Todo struct {
	Title       string
	Description string
	Completed   bool
}

func main() {
	http.HandleFunc("/", getTados)
	http.HandleFunc("/add-todo", addTodo)

	log.Println("Starting server on http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Error starting server: ", err)
	}

}

func getTados(w http.ResponseWriter, r *http.Request) {

	tmpl := template.Must(template.ParseFiles("./template/index.html"))
	todos := []Todo{
		{Title: "Task 1", Description: "Description 1", Completed: false},
		{Title: "Task 2", Description: "Description 2", Completed: true},
		{Title: "Task 3", Description: "Description 3", Completed: false},
	}
	tmpl.Execute(w, todos)
}
func addTodo(w http.ResponseWriter, r *http.Request) {

	title := r.PostFormValue("title")
	description := r.PostFormValue("description")
	completed := r.PostFormValue("completed") == "on"
	templ := template.Must(template.ParseFiles("./template/index.html"))
	templ.ExecuteTemplate(w, "todo-list-element", Todo{Title: title, Description: description, Completed: completed})

}
