package main

import (
    "fmt"
    "log"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "This is 2ver (branch feature1). Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
    http.HandleFunc("/", handler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}

func StartServer(){
    http.HandleFunc("/", handler
    log.Fatal(http.ListenAndServe(":8080", nil))
}

