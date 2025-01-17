package main

import (
  "fmt"
  "html"
  "log"
  "net/http"

  "github.com/gorilla/mux"

  "time"
  "encoding/json"
)

type Todo struct {
  // json doesn't like upper case keys        
  Name string `json:"name"`
  Completed bool `json:"completed"`
  Due time.Time `json:"due"`
}

type Todos []Todo

func main() {
  router := mux.NewRouter().StrictSlash(true)

  // waits for a request agaisnt index, if it receives it, calls Index()
  router.HandleFunc("/", Index)
  router.HandleFunc("/todos", TodoIndex)
  // this allows todoIds to be passed which allow us to respond with the proper record
  router.HandleFunc("/todos/{todoId}", TodoShow)


  log.Fatal(http.ListenAndServe(":8080", router))
}

func Index(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func TodoIndex(w http.ResponseWriter, r *http.Request) {
  todos := Todos{
    Todo{Name: "Write presentation"},
    Todo{Name: "Host meetup"},
  }

  json.NewEncoder(w).Encode(todos)
}

func TodoShow(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  todoId := vars["todoId"]
  fmt.Fprintln(w, "Todo show:", todoId)
}

