package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func hello(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("view/index.html")
	if err != nil {
		fmt.Println("Failed to parse index.html template.", err)
		return
	}

	tmpl.Execute(w, "World")
}

func main() {
	http.HandleFunc("/", hello)
	err := http.ListenAndServe("127.0.0.1:9300", nil)
	if err != nil {
		fmt.Println("Failed to start HTTP server, error:", err)
	}
}
