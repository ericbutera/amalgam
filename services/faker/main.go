package main

import (
	"encoding/xml"
	"log"
	"net/http"
	"strconv"

	"github.com/ericbutera/amalgam/internal/http/server"
	"github.com/ericbutera/amalgam/internal/test/faker/rss"
	"github.com/google/uuid"
	"github.com/samber/lo"
)

const DefaultItemCount int = 25

func main() {
	http.HandleFunc("/feed/", rssHandler)
	server := lo.Must(server.New())
	log.Fatal(server.ListenAndServe())
}

func rssHandler(w http.ResponseWriter, r *http.Request) {
	feedId := r.URL.Path[len("/feed/"):]
	a, _ := strconv.Atoi(r.URL.Query().Get("count"))
	count := lo.CoalesceOrEmpty(a, DefaultItemCount)

	if _, err := uuid.Parse(feedId); err != nil {
		http.Error(w, "Invalid UUID", http.StatusBadRequest)
		return
	}

	rss, err := rss.Generate(feedId, count)
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
