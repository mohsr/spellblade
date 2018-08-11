package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
)

func main() {
    /* Setup static content serving. */
    fs := http.FileServer(http.Dir("client"))

    /* Listen on env:PORT. */
    err := http.ListenAndServe(":" + os.Getenv("PORT"), nil)
    if err != nil {
        log.Fatal("ListenAndServe error: ", err)
    }
}
