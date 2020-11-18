package disk

import (
  "fmt"
  "io/ioutil"
  "net/http"
  "html/template"
  "os"
  "strings"
  "gopkg.in/yaml.v2"

  "github.com/ChrisPHP/MarbleFileServer/cookies"
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

//Repetitive code from the postgre.go file
//#########################################################
type Drives struct {
  Storage string
  Label string
}

type config struct {
  Name string
  password string
  Host string
  Port int
  Sqluser string
  Sqlpassword string
  Dbname string
  Drives []Drives
}

func (c *config) YamlReader() *config {
  file, err := ioutil.ReadFile("config.yaml")
  if err != nil {
    fmt.Println(err)
  }

  err = yaml.Unmarshal([]byte(file), c)
  if err != nil {
    fmt.Println(err)
  }

  return c
}
//#########################################################


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
  if (cookies.AuthCheck(w, r) != true) {
    http.Error(w, "Forbidden", http.StatusForbidden)
    return
  }


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

  if _, err = os.Stat(r.FormValue("dirs")); os.IsNotExist(err) {
    tmpl := template.Must(template.ParseFiles("./static/err.html"))
    tmpl.Execute(w, err)
    fmt.Println(err)
    return
  }

  var dir Dir
  var conf config
  conf.YamlReader()


  for _, elem := range conf.Drives {
    if (r.FormValue("PrevDir") == "") {
      dir.PrevDir = r.FormValue("dirs")
    } else {
      var x string
      if (r.FormValue("dirs") != elem.Storage) {
        parts := strings.SplitAfter(r.FormValue("dirs"), "/")
        for i := 0; i < len(parts) - 2; i++ {
          x += parts[i]
        }
        dir.PrevDir = x
      } else {
        dir.PrevDir = r.FormValue("dirs")
      }
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
