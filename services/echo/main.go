package main

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/ericbutera/amalgam/internal/http/server"
	"github.com/samber/lo"
)

func webhookHandler(w http.ResponseWriter, r *http.Request) {
	slog.Info("Received webhook",
		"method", r.Method,
		"url", r.URL,
		"remote_addr", r.RemoteAddr,
		"user_agent", r.UserAgent(),
		"content_length", r.ContentLength,
	)

	var payload map[string]any

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	data, err := json.MarshalIndent(payload, "", "  ")
	if err != nil {
		slog.Error("Failed to marshal payload", "error", err)
		http.Error(w, "Error processing request", http.StatusInternalServerError)

		return
	}

	w.WriteHeader(http.StatusOK)
	lo.Must(fmt.Fprintf(w, "headers: %+v\n", r.Header))
	lo.Must(fmt.Fprintf(w, "payload: %s\n", string(data)))

	if _, err := fmt.Fprintln(w, "Alert received"); err != nil {
		slog.Error("Failed to write response", "error", err)
	}
}

func main() {
	http.HandleFunc("/webhook", webhookHandler)

	port := "8080"
	slog.Info("Webhook receiver is running", "port", port)

	server := lo.Must(server.New())

	err := server.ListenAndServe()
	if err != nil {
		slog.Error("Failed to start server", "error", err)
	}
}
