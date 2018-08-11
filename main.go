package main

import (
    "log"
    "net/http"
    "os"
)

func main() {
    /* Grab environment variables. */
    port := os.Getenv("PORT")
    if port == "" {
        port = "80"
    }

    /* Setup static content serving. */
    http.Handle("/", http.FileServer(http.Dir("client")))

    /* Listen on a port. */
    err := http.ListenAndServe(":" + port, nil)
    if err != nil {
        log.Fatal("Error: ", err)
    }
}
