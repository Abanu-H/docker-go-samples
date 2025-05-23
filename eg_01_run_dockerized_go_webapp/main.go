package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Struct with exported field
type Goals struct {
	UserGoal string `json:"user_goal"`
}

// Global variable to store the current goal
var userGoal string

func getTemplate() string {
	return `<!DOCTYPE html>
<html>
	<head>
		<link rel="stylesheet" href="/styles.css">
	</head>
	<body>
		<section>
			<h2>My Course Goal</h2>
			<h3>{{.UserGoal}}</h3>
		</section>
		<form action="/store-goal" method="POST">
			<div class="form-control">
				<label>Course Goal</label>
				<input type="text" name="goal">
			</div>
			<button>Set Course Goal</button>
		</form>
	</body>
</html>`
}

func main() {
	r := mux.NewRouter()

	// Basic GET Api to list the goal
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.New("goal").Parse(getTemplate()))

		goal := userGoal
		if goal == "" {
			goal = "Learn Docker!"
		}

		data := Goals{UserGoal: goal}
		tmpl.Execute(w, data)
	}).Methods("GET")

	// BAisc POST Api to Update the goal
	r.HandleFunc("/store-goal", func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		userGoal = r.FormValue("goal")

		// Redirect to GET /
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}).Methods("POST")

	// Serve static files (styles.css) from public/
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./public/")))

	//Port Assigment
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
