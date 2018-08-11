package main

import (
    "encoding/json"
    "cloud.google.com/go/firestore"
    "log"
    "net/http"
    "os"
)

type Result struct {
    Text  string
    Color string
}

func main() {
    /* Grab environment variables. */
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    /* Serve static content. */
    files := http.FileServer(http.Dir("client"))
    http.Handle("/", files)

    /* POST a text command. */
    http.HandleFunc("/cmd", func(w http.ResponseWriter, r *http.Request) {
        /* Enforce POST method. */
        if r.Method != "POST" {
            http.Error(w, "Bad request method.", 405)
            return
        }

        /* Create a response object to send back as JSON. */
        res := Result{Text: "Sorry, invalid command!", Color: "Red"}

        /* Parse the command and formulate the appropriate response. */
        if txt := r.FormValue("txt"); txt != "" {
            res.Text  = "TOBEPARSED"
            res.Color = "White"
        } 

        /* Create and write the JSON. */
        js, err := json.Marshal(res)
        if err != nil {
            log.Fatal("Error: ", err)
        }
        w.Header().Set("Content-Type", "application/json")
        w.Write(js)
    })

    /* Listen on a port. */
    err := http.ListenAndServe(":" + port, nil)
    if err != nil {
        log.Fatal("Error: ", err)
    }
}
