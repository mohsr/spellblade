package main

import (
    "log"
    "net/http"
    "os"
)

func main() {
    /* Grab environment variables. */
    port := os.Getenv("PORT")
    if port != nil {
        port = "80"
    }

    /* Setup static content serving. */
    fs := http.FileServer(http.Dir("client"))

    /* Listen on env:PORT. */
    err := http.ListenAndServe(":" + port, nil)
    if err != nil {
        log.Fatal("Error: ", err)
    }
}
