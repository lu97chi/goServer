package main

import (
	"net/http"
)

func routes() {
	http.HandleFunc("/addTodo", addTodo)

	http.HandleFunc("/deleteTodo", deleteTodo)

	http.HandleFunc("/updateTodo", updateTodo)

	http.HandleFunc("/getTodo", getTddo)
}
