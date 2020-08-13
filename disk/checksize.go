package disk

import (
  "fmt"
  "io/ioutil"
  "net/http"
  "html/template"
  "os"
  "strings"
)

type Content struct {
  FileDir string
  File string
  FoldDir string
  Fold string
}

type Dir struct {
  Contents []Content
  CurDir string
  PrevDir string
  TheDir string
}

type Redirect struct {
  Result string
  Fold string
}

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
  //Delete file Handler
  err := os.Remove(r.FormValue("DelFile"))
  if err != nil {
    fmt.Println(err)
    return
  }
  DirHandler(w, r)
}

//Serve file for download or viewing
func DownloadHandler(w http.ResponseWriter, r *http.Request) {
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

  if (r.FormValue("PrevDir") == "") {
    dir.PrevDir = "/Marble1/"
  } else {
    var x string
    if (r.FormValue("dirs") != "/Marble1/") {
      parts := strings.SplitAfter(r.FormValue("dirs"), "/")
      for i := 0; i < len(parts) - 2; i++ {
        x += parts[i]
      }
      dir.PrevDir = x
    } else {
      dir.PrevDir = r.FormValue("dirs")
    }
  }

  dir.CurDir = r.FormValue("dirs")

  //Add the file names to the struct
  for _, file := range files {
    if (file.IsDir() == true) {
      dir.Contents = append(dir.Contents, Content{
        FoldDir: r.FormValue("dirs") + file.Name(),
        Fold: file.Name(),
      })
    } else {
      dir.Contents = append(dir.Contents, Content{
        FileDir: r.FormValue("dirs") + file.Name(),
        File: file.Name(),
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

  FilePath := r.FormValue("dirs") + "/" + r.FormValue("newDir")

  err := os.Mkdir(FilePath, 0755)
  if (err != nil) {
    fmt.Fprintf(w, "failed to create directory")
    fmt.Println(err)
    return
  }

  DirHandler(w, r)
}
