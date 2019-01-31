package main

import (
    "fmt"
    "net/http"

    "github.com/prometheus/client_golang/prometheus"
    "github.com/urfave/negroni"
    "github.com/zbindenren/negroni-prometheus"
)

func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
        fmt.Fprintf(w, "Welcome to the home page!")
    })
    mux.Handle("/metrics", prometheus.Handler())

    n := negroni.New()
    n.Use(negroni.NewLogger())
    n.Use(negroniprometheus.NewMiddleware("negroni"))
    n.UseHandler(mux)

    http.ListenAndServe(":8080", n)
}
