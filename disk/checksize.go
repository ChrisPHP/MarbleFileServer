package disk

import (
  "fmt"
  "io/ioutil"
  "net/http"
  "html/template"
  "os"
)

type Content struct {
  File string
}

type Dir struct {
  Contents []Content
}

func DirHandler(w http.ResponseWriter, r *http.Request) {
  tmpl := template.Must(template.ParseFiles("./static/test.html"))

  if r.Method != http.MethodPost {
    tmpl.Execute(w, nil)
    return
  }

  //Read the directory
  files, err := ioutil.ReadDir(r.FormValue("dirs"))
  if err != nil {
    fmt.Println(err)
  }

  var dir Dir

  //Add the file names to the struct
  for _, file := range files {
    dir.Contents = append(dir.Contents, Content{
      File: file.Name(),
    })
    fmt.Println(dir)
  }

  //Items found in directory is listed onto the template
  err = tmpl.Execute(w, dir)
    if err != nil {
      fmt.Println(err)
    }
}

func CreateDirHandler(w http.ResponseWriter, r *http.Request) {
  if r.Method != http.MethodPost {
    fmt.Fprintf(w, "No Directory name given")
    return
  }

  err := os.Mkdir(r.FormValue("newDir") ,0755)
  if (err != nil) {
    fmt.Fprintf(w, "failed to create directory")
    fmt.Println(err)
    return
  }

  fmt.Fprintf(w, "Directory has been made")
}

func ShowDir() {
  files, err := ioutil.ReadDir("./images")
  if err != nil {
    fmt.Println(err)
  }

  var dir Dir

  for _, file := range files {
    dir.Contents = append(dir.Contents, Content{
      File: file.Name(),
    })
    fmt.Println(dir)
  }
}
