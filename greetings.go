// greetings.go
package main

import (
    "fmt"
    "net/http"
    "os"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, welcome to the Heroku-deployed Go app!")
}

func main() {
    http.HandleFunc("/", handler)

    // Use the PORT environment variable on Heroku
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080" // Default port if not set
    }

    fmt.Printf("Listening on port %s...\n", port)
    if err := http.ListenAndServe(":"+port, nil); err != nil {
        fmt.Println("Error starting server:", err)
    }
}

