package wotd

import (
	"fmt"
	"net/http"
	"time"
)

// Handler represents an HTTP handler for the word generator.
type Handler struct {
	Generator interface {
		Generate(word string) (string, error)
	}

	Now func() time.Time
}

// NewHandler returns an instance of Handler with default settings.
func NewHandler() *Handler {
	return &Handler{Now: time.Now}
}

// ServeHTTP writes an HTML page with the word of the day.
func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Retrieve the current day of the week.
	day := h.Now().Format("Monday")

	// Generate a phrase for the day.
	phrase, err := h.Generator.Generate(day)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Write phrase to output.
	fmt.Fprintln(w, "<html>")
	fmt.Fprintln(w, "<style> h1 { font-size: 84px; } </style>")
	fmt.Fprintln(w, "<body>")
	fmt.Fprintf(w, "<center><h1>Have a<br>%s!</h1></center>\n", phrase)
	fmt.Fprintln(w, "</body>")
	fmt.Fprintln(w, "</html>")
}
