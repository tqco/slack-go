package hello

import (
    "fmt"
    "net/http"
)

func init() {
    http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
  //Read the Request Parameter "command"
  command := r.FormValue("command")

  // TODO: Other checks for tokens/username/etc

  if command == "/menu" {
   fmt.Fprint(w,"See the full club menu at http://go.tq.co/menu. Nom!")
  } else {
    fmt.Fprint(w,"Error")
  }
}
