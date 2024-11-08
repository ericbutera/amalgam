package main

import (
	"encoding/xml"
	"log"
	"net/http"

	"github.com/ericbutera/amalgam/internal/http/server"
	"github.com/ericbutera/amalgam/internal/test/faker/rss"
	"github.com/google/uuid"
	"github.com/samber/lo"
)

func main() {
	http.HandleFunc("/feed/", rssHandler)
	server := lo.Must(server.New())
	log.Fatal(server.ListenAndServe())
}

func rssHandler(w http.ResponseWriter, r *http.Request) {
	feedId := r.URL.Path[len("/feed/"):]
	if _, err := uuid.Parse(feedId); err != nil {
		http.Error(w, "Invalid UUID", http.StatusBadRequest)
		return
	}

	rss, err := rss.Generate(feedId)
	if err != nil {
		http.Error(w, "Error generating feed", http.StatusInternalServerError)
		return
	}

	output, err := xml.MarshalIndent(rss, "", "  ")
	if err != nil {
		http.Error(w, "Error generating XML", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/rss+xml")
	lo.Must(w.Write(output))
}
