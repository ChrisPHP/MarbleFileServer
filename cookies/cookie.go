package cookies

import (
  "net/http"
  "fmt"

  "github.com/gorilla/sessions"
)

var (
  Key = []byte(make([]byte, 32))
  store = sessions.NewCookieStore(Key)
)

func CookieUsername(w http.ResponseWriter, r *http.Request) string {
  session, _ := store.Get(r, "Marble-Cookie")

  Name := fmt.Sprintf("%v", session.Values["UserName"])

  return Name
}

func AuthCheck(w http.ResponseWriter, r *http.Request) bool {
  session, _ := store.Get(r, "Marble-Cookie")

  if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
    return false
  }
  return true
}

func CookieCreate(w http.ResponseWriter, r *http.Request, Name string) {
  session, _ := store.Get(r, "Marble-Cookie")

  session.Values["authenticated"] = true
  session.Values["UserName"] = Name
  session.Save(r, w)
}

func CookieRemove(w http.ResponseWriter, r *http.Request) {
  session, _ := store.Get(r, "Marble-Cookie")

  session.Values["authenticated"] = false
  session.Save(r, w)
}
