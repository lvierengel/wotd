package wotd

import (
	"errors"
	"math/rand"
	"strings"
)

//go:generate go-bindata -o assets/assets.go -pkg assets -ignore assets.go -prefix assets/ assets

var (
	// ErrAdjectiveNotFound is returned when an adjective with the given first letter cannot be found.
	ErrAdjectiveNotFound = errors.New("adjective not found")

	// ErrTooShort is returned when the input to Generate() is empty.
	ErrTooShort = errors.New("too short")
)

// Generator represents a word generator that returns a phrase based on a given word.
type Generator struct {
	words map[string][]string // words by first letter
}

// NewGenerator returns a new instance of Generator with the given word set.
func NewGenerator(words []string) *Generator {
	// Group words by first letter.
	m := make(map[string][]string)
	for _, w := range words {
		// Skip blank words.
		if len(w) == 0 {
			continue
		}

		// Append word by initial letter.
		letter := string(w[0])
		m[letter] = append(m[letter], w)
	}

	return &Generator{words: m}
}

// Generate returns a phrase based on word.
func (g *Generator) Generate(word string) (string, error) {
	// Return an error if a blank string is passed in.
	if len(word) == 0 {
		return "", ErrTooShort
	}

	// Retrieve a list of adjectives starting with the first letter.
	a := g.words[strings.ToLower(string(word[0]))]
	if len(a) == 0 {
		return "", ErrAdjectiveNotFound
	}

	// Randomly choose an adjective.
	adj := a[rand.Intn(len(a))]
	adj = strings.ToUpper(string(adj[0])) + adj[1:]

	// Return the concatentated phrase.
	return adj + " " + word, nil
}
