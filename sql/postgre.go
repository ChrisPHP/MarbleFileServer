package sql

import (
    "database/sql"
    "fmt"
    "net/http"
    "html/template"

    "github.com/ChrisPHP/MarbleFileServer/cookies"
    "golang.org/x/crypto/bcrypt"
    _ "github.com/lib/pq"
)

const (
  host = "localhost"
  port = 5432
  sqluser = "postgres"
  sqlpassword = "SQL-PASSWORD-HERE"
  dbname  = "DATABASENAME-HERE"
)

type users struct {
  Name string
  password string
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

  var User users

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
  psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+ "password=%s dbname=%s sslmode=disable", host, port, sqluser, sqlpassword, dbname)

  db, err := sql.Open("postgres", psqlInfo)
  if (err != nil) {
    fmt.Println(err)
  }
  defer db.Close()

  err = db.Ping()
  if (err != nil) {
    fmt.Println(err)
  }

  fmt.Println("Successfully connected!")

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

  psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+ "password=%s dbname=%s sslmode=disable", host, port, sqluser, sqlpassword, dbname)

  db, err := sql.Open("postgres", psqlInfo)
  if (err != nil) {
    fmt.Println(err)
  }
  defer db.Close()

  err = db.Ping()
  if (err != nil) {
    fmt.Println(err)
  }

  fmt.Println("Successfully connected!")

  var User users

  QueryString := "SELECT password FROM users WHERE username = $1"

  row := db.QueryRow(QueryString, r.FormValue("username"))

  err = row.Scan(&User.password)
  if err != nil {
    if err == sql.ErrNoRows {
      fmt.Println("No rows were returned")
      return
    }
    fmt.Println(err)
    return
  }

  if (checkHash(r.FormValue("password"), User.password) != true) {
    fmt.Println("Passwords do not match")
    return
  }

  cookies.CookieCreate(w, r, r.FormValue("username"))

  User.Name = r.FormValue("username")

  err = tmpl.Execute(w, User)
    if err != nil {
      fmt.Println(err)
    }
}
