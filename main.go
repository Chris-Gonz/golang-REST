package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/justinas/alice"
	_ "github.com/lib/pq"
)

type RouteResponse struct {
	Message string `json:"message"`
	ID      string `json:"id,omitempty"`
}

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dbPassword := os.Getenv("DB_PASSWORD")

	connectInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)

	db, err := sql.Open("postgres", connectInfo)

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Since Open can return just the validation of the arguments, best practice to make sure a connection is made
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("Connection to db made")

	routerMux := http.NewServeMux()

	routerMux.Handle("/register", alice.New(loggingMiddleware).ThenFunc(register))
	routerMux.Handle("/login", alice.New(loggingMiddleware).ThenFunc(login))
	routerMux.Handle("/projects", alice.New(loggingMiddleware).ThenFunc(getProjects))     // this can also call create{POST}
	routerMux.Handle("/projects/{id}", alice.New(loggingMiddleware).ThenFunc(getProject)) //this can also call Update{GET}
	routerMux.Handle("/deleteProject/{id}", alice.New(loggingMiddleware).ThenFunc(deleteProject))

	log.Fatal(http.ListenAndServe(":5000", routerMux))

}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
	})
}

//Good idea to plan out before hand what your API will do.
//In this project there will be simple CRUD api

// register
func register(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not Allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(RouteResponse{"Hello from register!", ""})
}

// login
func login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not Allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(RouteResponse{"Hello from login!", ""})
}

// createProject
func createProject(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not Allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(RouteResponse{"Hello from createProject!", ""})
}

// updateProject
func updateProject(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method not Allowed", http.StatusMethodNotAllowed)
		return
	}

	//Get the id field from route response, parse path and get id
	id := r.URL.Path[len("/projects/"):]

	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(RouteResponse{"Hello from update!", id})
}

// getProjects - get all projects and create a new project
func getProjects(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodPost:
		createProject(w, r)
	case http.MethodGet:
		w.Header().Set("Content-type", "application/json")
		json.NewEncoder(w).Encode(RouteResponse{"Hello from getProjects!", ""})
	default:
		if r.Method != http.MethodGet {
			http.Error(w, "Method not Allowed", http.StatusMethodNotAllowed)
			return
		}
	}

}

// getProject
func getProject(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodPut:
		updateProject(w, r)
	case http.MethodGet:
		w.Header().Set("Content-type", "application/json")
		json.NewEncoder(w).Encode(RouteResponse{"Hello from getProject!", ""})
	default:
		http.Error(w, "Method not Allowed", http.StatusMethodNotAllowed)
	}

}

// deleteProject
func deleteProject(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not Allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(RouteResponse{"Hello from deleteProject!", ""})
}
