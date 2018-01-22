package main

import (
  "net/http"

  "github.com/gorilla/mux"
)

func NewRouter() *mux.Router {

  router := mux.NewRouter().StrictSlash(true)
  for _, route := range routes {
    var handler http.Handler

    handler = route.HandlerFunc
    handler = Logger(handler, route.Name) // wraps the handler with a loging and timing capabilites

    router.
      Methods(route.Method).
      Path(route.Pattern).
      Name(route.Name).
      Handler(handler)
  }

  return router
}
