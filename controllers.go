package main

import (
	"fmt"
	"net/http"
)

func addTodo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the add route")
}

func deleteTodo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the delete route")
}

func updateTodo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the update route")
}

func getTddo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the get route")
}
