package main

import (
    "encoding/json"
    "log"
    "net/http"
    "os"

    firebase "firebase.google.com/go"
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

    /* Setup Cloud Firestore. */
    conf := &firebase.Config{ProjectID: "firebase-go"}
    app, err := firebase.NewApp(ctx, conf)
    if err != nil {
        log.Fatalln(err)
    }
    client, err := app.Firestore(ctx)
    if err != nil {
        log.Fatalln(err)
    }
    defer client.Close()

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
            http.Error(w, "Could not form JSON.", 500)
        }
        w.Header().Set("Content-Type", "application/json")
        w.Write(js)
    })

    /* Listen on a port. */
    err := http.ListenAndServe(":" + port, nil)
    if err != nil {
        log.Fatalln(err)
    }
}
