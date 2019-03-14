package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello, world!")
    })
    http.HandleFunc("qa.example.com/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello, improved world!")
    })
    http.ListenAndServe(":8080", nil)
}