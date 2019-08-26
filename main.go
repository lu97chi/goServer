package main

import (
	"fmt"
	"net/http"
)

// type Trainer struct {
// 	Name string
// 	Age  int
// 	City string
// }

func main() {

	dbConection()
	routes()

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	fmt.Print("Connected to port 3001")

	http.ListenAndServe(":3001", nil)
}
