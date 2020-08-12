package uploads

import (
  "net/http"
  "io/ioutil"
  "mime/multipart"
  "fmt"

  "github.com/ChrisPHP/MarbleFileServer/disk"
)

func UploadHandler(w http.ResponseWriter, r *http.Request) {
  //If the URL path does not have "/upload" throw a 404 error
  if r.URL.Path != "/upload" {
    http.Error(w, "404 not found.", http.StatusNotFound)
    return
  }

  r.ParseMultipartForm(32 << 20)
  files := r.MultipartForm.File["myFile"]

  for _, handler := range files {
    fmt.Println(handler.Filename)
    file, err := handler.Open()
    defer file.Close()
    if err != nil {
      fmt.Println(err)
      return
    }

    SaveFile(w, file, handler, r)
  }

  disk.DirHandler(w, r)
}

func SaveFile(w http.ResponseWriter, file multipart.File, handler *multipart.FileHeader, r *http.Request) {
  //Read all the users data
  data, err := ioutil.ReadAll(file)
  if err != nil {
    fmt.Fprintf(w, "%v", err)
    return
  }

  //Write the data to the servers storage
  err = ioutil.WriteFile(r.FormValue("dirs")+handler.Filename, data, 0666)
  if err != nil {
    fmt.Fprintf(w, "%v", err)
    return
  }
}
