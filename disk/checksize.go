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
  Fold string
  Dime string
}

type Dir struct {
  Contents []Content
  CurDir string
}

func DownloadHandler(w http.ResponseWriter, r *http.Request) {
  fmt.Println(r.FormValue("FoldFile"))
  http.ServeFile(w, r, r.FormValue("FoldFile"))
}

func DirHandler(w http.ResponseWriter, r *http.Request) {
  tmpl := template.Must(template.ParseFiles("./static/ViewDir.html"))


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

  dir.CurDir = r.FormValue("dirs")

  //Add the file names to the struct
  for _, file := range files {
    if (file.IsDir() == true) {
      dir.Contents = append(dir.Contents, Content{
        Fold: r.FormValue("dirs") + file.Name(),
      })
    } else {
      dir.Contents = append(dir.Contents, Content{
        File: r.FormValue("dirs") + file.Name(),
      })
    }
  }

  //Items found in directory is listed onto the template
  err = tmpl.Execute(w, dir)
    if err != nil {
      fmt.Println(err)
    }
}


//Create a directory
func CreateDirHandler(w http.ResponseWriter, r *http.Request) {
  if r.Method != http.MethodPost {
    fmt.Fprintf(w, "No Directory name given")
    return
  }

  FilePath := r.FormValue("Fold") + "/" + r.FormValue("newDir")

  err := os.Mkdir(FilePath, 0755)
  if (err != nil) {
    fmt.Fprintf(w, "failed to create directory")
    fmt.Println(err)
    return
  }

  fmt.Fprintf(w, "Directory has been made")
}
