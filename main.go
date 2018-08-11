package main

import (
    "encoding/json"
    "github.com/gorilla/mux"
    "log"
    "net/http"
    "os"
)

type result struct {
    text  string
    color string
}

func main() {
    /* Grab environment variables. */
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    /* Setup server router. */
    r := mux.NewRouter()

    /* POST a text command. */
    r.HandleFunc("/cmd", func(w http.ResponseWriter, r *http.Request) {
        /* Create a response object to send back as JSON. */
        res := result{text: "Sorry", color: "Red"}

        /* Parse the command and formulate the appropriate response. */
        if txt := r.FormValue("txt"); txt != "" {
            res.text  = "Sorry, invalid command!"
            res.color = "Red"
        } 

        js, err := json.Marshal(res)
        if err != nil {
            log.Fatal("Error: ", err)
        }

        w.Header().Set("Content-Type", "application/json")
        w.Write(js)
    }).Methods("POST")

    r.PathPrefix("/").Handler(http.FileServer(http.Dir("client")))

    /* Listen on a port. */
    err := http.ListenAndServe(":" + port, r)
    if err != nil {
        log.Fatal("Error: ", err)
    }
}
