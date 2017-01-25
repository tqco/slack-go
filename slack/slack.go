package slack

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
    fmt.Fprint(w,"Check out the club menu at http://go.tq.co/menu. Nom!")
  } else if command == "/manual" {
    fmt.Fprint(w,"Check out the manual at http://go.tq.co/manual. Nom!")
  } else if command == "/events" {
    fmt.Fprint(w,"Check out our events calendar at http://tq.co/events or subscribe to http://go.tq.co/events. Nom!")
  } else if command == "/oo" {
    fmt.Fprint(w,"Want help or advice from TQ? Come by during our office hours via http://go.tq.co/oo")
  } else if command == "/errors" {
    fmt.Fprint(w,"Spotted something wrong at TQ? File an error by sending an email to errors@tq.co.")
  } else {
    fmt.Fprint(w,"")
  }
}
