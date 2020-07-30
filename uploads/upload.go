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

  //Recieve users submitted files
  file, handler, err := r.FormFile("myFile")
  if err != nil {
    fmt.Println("Error to recieve files")
    fmt.Println(err)
    return
  }
  defer file.Close()

  //prints to console some file details
  fmt.Printf("Uploaded File: %+v\n", handler.Filename)
  fmt.Printf("File Size: %+v\n", handler.Size)
  fmt.Printf("MIME header: %+v\n", handler.Header)

  SaveFile(w, file, handler, r)
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
