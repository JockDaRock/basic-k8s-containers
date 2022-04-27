package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/", WhatsUpBeautifulPeople)
    http.ListenAndServe(":8080", nil)
}

func WhatsUpBeautifulPeople(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Do you know the shoulder touch?,  HEY! %s! **ZAP**", r.URL.Path[1:])
}