package sql

import (
    "database/sql"
    "fmt"
    "net/http"
    "html/template"
    "io/ioutil"
    "gopkg.in/yaml.v2"

    "github.com/ChrisPHP/MarbleFileServer/cookies"
    "golang.org/x/crypto/bcrypt"
    _ "github.com/lib/pq"
)

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

func hasher(pass string) (string) {
  bytes, err := bcrypt.GenerateFromPassword([]byte(pass), 14)
  if (err != nil) {
    fmt.Println(err)
    return "failed to hash"
  }
  return string(bytes)
}

func checkHash(pass, hash string) bool {
  err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pass))
  if (err != nil) {
    fmt.Println(err)
  }
  return err == nil
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
  if (cookies.AuthCheck(w, r) != true) {
    tmpl := template.Must(template.ParseFiles("./static/index.html"))

    err := tmpl.Execute(w, nil)
      if err != nil {
        fmt.Println(err)
      }
      return
  }

  tmpl := template.Must(template.ParseFiles("./static/dirs.html"))

  var User config

  User.Name = cookies.CookieUsername(w, r)

  err := tmpl.Execute(w, User)
    if err != nil {
      fmt.Println(err)
    }
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
  cookies.CookieRemove(w, r)
  IndexHandler(w, r)
}

func ChangeHandler(w http.ResponseWriter, r *http.Request) {
  var c config
  c.YamlReader()

  psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+ "password=%s dbname=%s sslmode=disable", c.Host, c.Port, c.Sqluser, c.Sqlpassword, c.Dbname)

  db, err := sql.Open("postgres", psqlInfo)
  if (err != nil) {
    fmt.Println(err)
  }
  defer db.Close()

  err = db.Ping()
  if (err != nil) {
    fmt.Println(err)
  }

  hash := hasher(r.FormValue("password"))

  QueryString := "UPDATE users SET password = $2 WHERE username = $1;"
  _, err = db.Exec(QueryString, cookies.CookieUsername(w, r), hash)
  if err != nil {
    fmt.Println(err)
  }

  LogoutHandler(w, r)
}

func SigninHandler(w http.ResponseWriter, r *http.Request) {
  tmpl := template.Must(template.ParseFiles("./static/dirs.html"))

  var c config
  c.YamlReader()

  psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+ "password=%s dbname=%s sslmode=disable", c.Host, c.Port, c.Sqluser, c.Sqlpassword, c.Dbname)

  db, err := sql.Open("postgres", psqlInfo)
  if (err != nil) {
    fmt.Println(err)
  }
  defer db.Close()

  err = db.Ping()
  if (err != nil) {
    fmt.Println(err)
  }

  QueryString := "SELECT password FROM users WHERE username = $1"

  row := db.QueryRow(QueryString, r.FormValue("username"))

  err = row.Scan(&c.password)
  if err != nil {
    if err == sql.ErrNoRows {
      fmt.Println("No rows were returned")
      return
    }
    fmt.Println(err)
    return
  }

  if (checkHash(r.FormValue("password"), c.password) != true) {
    fmt.Println("Passwords do not match")
    return
  }

  cookies.CookieCreate(w, r, r.FormValue("username"))

  c.Name = r.FormValue("username")

  err = tmpl.Execute(w, c)
    if err != nil {
      fmt.Println(err)
    }
}
