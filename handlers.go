package main

import (
  "encoding/json"
  "fmt"
  "net/http"
  "io"
  "io/ioutil"


  "github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintln(w, "Welcome!")
}

func TodoIndex(w http.ResponseWriter, r *http.Request) {
  // tell the client to expect json and explicitly set the status code
  // go will no longer guess (it could have been wrong)
  w.Header().Set("Content-Type", "application/json; charset=UTF-8")
  w.WriteHeader(http.StatusOK)
  if err := json.NewEncoder(w).Encode(todos); err != nil {
    panic(err)
  }
}

func TodoShow(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  todoId := vars["todoId"]
  fmt.Fprintln(w, "Todo show:", todoId)
}

func TodoCreate(w http.ResponseWriter, r *http.Request) {
  var todo Todo
  // using io.LimitReader to protect against malicious attacks (say 500GB of json received)
  body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048567))
  if err != nil {
    panic(err)
  }
  if err := r.Body.Close(); err != nil {
    panic(err)
  }
  // unmarshall it into the Todo struct
  if err := json.Unmarshal(body, &todo);err != nil {
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(422)  // unprocessable entity
    if err := json.NewEncoder(w).Encode(err); err != nil {
      panic(err)
    }
    panic(err)
  }

  t := RepoCreateTodo(todo)
  w.Header().Set("Content-Type", "application/json; charset=UTF-8")
  w.WriteHeader(http.StatusCreated)   // status code:201
  if err := json.NewEncoder(w).Encode(t); err != nil {
    panic(err)
  }
}
