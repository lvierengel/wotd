package wotd

import (
	"fmt"
	"net/http"
	"strings"
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

	// Generate an adjective for the day.
	adj, err := h.Generator.Generate(day)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Write phrase to output.
	fmt.Fprintln(w, "<html>")
	fmt.Fprintln(w, "<style> h1 { font-size: 84px; }")
	fmt.Fprintln(w, "h1 { font-size: 84px; }")
	fmt.Fprintln(w, "a { color: black; text-decoration: none; }")
	fmt.Fprintln(w, "a:hover { text-decoration: underline; }")
	fmt.Fprintln(w, "</style>")
	fmt.Fprintln(w, "<body>")
	fmt.Fprintln(w, "<center>")
	fmt.Fprintln(w, "  <h1>")
	fmt.Fprintf(w, "    <a href=\"http://dictionary.reference.com/browse/%s\" target=\"_blank\">\n", adj)
	fmt.Fprintf(w, "      Have a<br>%s %s!\n", initialCase(adj), day)
	fmt.Fprintf(w, "    </a>\n")
	fmt.Fprintln(w, "  </h1>")
	fmt.Fprintln(w, "</center>")
	fmt.Fprintln(w, "</body>")
	fmt.Fprintln(w, "</html>")
}

// initialCase returns s with an uppercase first letter and lower case remaining letters.
func initialCase(s string) string {
	if s == "" {
		return ""
	}
	return strings.ToUpper(string(s[0])) + strings.ToLower(s[1:])
}
