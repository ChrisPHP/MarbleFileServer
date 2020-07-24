package main

import (
  "net/http"
  "fmt"
  "log"

  "github.com/ChrisPHP/MarbleFileServer/uploads"
  "github.com/ChrisPHP/MarbleFileServer/disk"
)

type TheFile struct {
  Itm string
  Price int
}

type Dir struct {
  Title string
  MyFiles []TheFile
}

func setupRoutes() {
  fileServer := http.FileServer(http.Dir("./static"))
  http.Handle("/", fileServer)

  http.HandleFunc("/Create", disk.CreateDirHandler)
  http.HandleFunc("/upload", uploads.UploadHandler)
  http.HandleFunc("/view", disk.DirHandler)
  http.HandleFunc("/Download", disk.DownloadHandler)

  if err := http.ListenAndServe(":8080", nil); err != nil {
    log.Fatal(err)
  }
}

func main() {
  fmt.Println("Go file upload")
  setupRoutes()
}
