package main

import (
  "net/http"
  "fmt"
)

func uploadFile(w http.ResponseWriter, r *http.Request) {
  fmt.Fprint(w, "Test")
}

func setupRoutes() {
  http.HandleFunc("/upload", uploadFile)
  http.ListenAndServe(":8080", nil)
}

func main() {
  fmt.Println("Go file upload")
  setupRoutes()
}
