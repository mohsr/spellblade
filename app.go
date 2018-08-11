package main

import (
    "fmt"
    "log"
    "net/http"
)

func main() {
    /* Handle calls to /. */
    http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
        fmt.Fprint(res, "Hello world!")
    })

    /* Listen on port 8080. */
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        log.Fatal("ListenAndServe error: ", err)
    }
}