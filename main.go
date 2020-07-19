package main

import (
  "net/http"
  "fmt"
  "log"

  "github.com/ChrisPHP/MarbleFileServer/uploads"
)

func setupRoutes() {
  fileServer := http.FileServer(http.Dir("./static"))
  http.Handle("/", fileServer)
  http.HandleFunc("/upload", uploads.UploadHandler)
  if err := http.ListenAndServe(":8080", nil); err != nil {
    log.Fatal(err)
  }
}

func main() {
  fmt.Println("Go file upload")
  setupRoutes()
}
