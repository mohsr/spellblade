package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
)

func main() {
    /* Handle calls to /. */
    http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
        fmt.Fprint(res, "Hello world!")
    })

    /* Listen on PORT env variable. */
    err := http.ListenAndServe(":" + os.Getenv("PORT"), nil)
    if err != nil {
        log.Fatal("ListenAndServe error: ", err)
    }
}
