package main

import (
    "encoding/json"
    // "context"
    // "google.golang.org/api/option"
    "log"
    "net/http"
    "os"

    // firebase "firebase.google.com/go"
)

/*
 * Holds the response of parsing a text command.
 * Schema: {"Text":  The response text to send back.
 *          "Color": The color of the response text to display.}
 */
type Response struct {
    Text  string
    Color string
}

func main() {
    /* Grab environment variables. */
    PORT := os.Getenv("PORT")
    if PORT == "" {
        PORT = "8080"
    }

    /* Setup Cloud Firestore with service account. */
    // ctx := context.Background()
    // sa := option.WithCredentialsFile("spellblade-go-private-key.json")
    // app, err := firebase.NewApp(ctx, nil, sa)
    // if err != nil {
    //     log.Fatalln(err)
    // }
    // db, err := app.Firestore(ctx)
    // if err != nil {
    //     log.Fatalln(err)
    // }
    // defer db.Close()

    /* Serve static content. */
    files := http.FileServer(http.Dir("client"))
    http.Handle("/", files)

    /* 
     * POST a text command.
     * Route:  /cmd
     * Schema: {"txt": Text of the command entered.}
     */
    http.HandleFunc("/cmd", func(w http.ResponseWriter, r *http.Request) {
        /* Enforce POST method. */
        if r.Method != "POST" {
            http.Error(w, "Bad request method.", 405)
            return
        }

        /* Create a response object to send back as JSON. */
        res := Response{Text: "Sorry, invalid command!", Color: "Red"}

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
    err := http.ListenAndServe(":" + PORT, nil)
    if err != nil {
        log.Fatalln(err)
    }
}
