package main

import (
  "log"
  "net/http"
  "time"
)

// no built in logging function so we have to write our own
func Logger(inner http.Handler, name string) http.Handler {
  // passed our handler
  // return the handler wrapped with logging and timing functionality
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    start := time.Now()

    inner.ServeHTTP(w, r)

    log.Printf(
      "%s\t%s\t%s\t%s",
      r.Method,
      r.RequestURI,
      name,
      time.Since(start),
   )

  })
}
