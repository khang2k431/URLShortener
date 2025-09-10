package main

import (
    "fmt"
    "log"
    "net/http"
    "urlshortener/service"
)

var urlService = service.NewURLService()

func shortenHandler(w http.ResponseWriter, r *http.Request) {
    longURL := r.URL.Query().Get("url")
    if longURL == "" {
        http.Error(w, "Missing url parameter", http.StatusBadRequest)
        return
    }
    short := urlService.GenerateShortURL()
    urlService.SaveURL(short, longURL)
    fmt.Fprintf(w, "Short URL: http://localhost:8080/%s\n", short)
}

func redirectHandler(w http.ResponseWriter, r *http.Request) {
    short := r.URL.Path[1:]
    longURL, exists := urlService.GetURL(short)
    if !exists {
        http.NotFound(w, r)
        return
    }
    http.Redirect(w, r, longURL, http.StatusFound)
}

func statsHandler(w http.ResponseWriter, r *http.Request) {
    stats := urlService.Stats()
    for short, count := range stats {
        fmt.Fprintf(w, "%s -> visits: %d\n", short, count)
    }
}

func main() {
    http.HandleFunc("/shorten", shortenHandler)
    http.HandleFunc("/stats", statsHandler)
    http.HandleFunc("/", redirectHandler)

    fmt.Println("Server started at :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
