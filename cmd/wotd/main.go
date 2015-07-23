package main

import (
	"flag"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/benbjohnson/wotd"
	"github.com/benbjohnson/wotd/assets"
)

func main() {
	m := NewMain()
	if err := m.Run(os.Args[1:]...); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// Main represents the main program execution.
type Main struct {
	Stdout io.Writer
}

// NewMain returns a new instance of Main.
func NewMain() *Main {
	return &Main{
		Stdout: os.Stdout,
	}
}

// Run executes the program.
func (m *Main) Run(args ...string) error {
	// Seed the random number generator.
	rand.Seed(time.Now().UnixNano())

	// Parse command line flags.
	opt, err := m.ParseFlags(args)
	if err != nil {
		return err
	}

	// Create word generator.
	data := assets.MustAsset("words")
	words := strings.Split(string(data), "\n")

	// Instantiate a new generator.
	g := wotd.NewGenerator(words)

	// Start HTTP server.
	h := wotd.NewHandler()
	h.Generator = g

	fmt.Fprintf(m.Stdout, "Listening on %s\n", opt.Addr)
	return http.ListenAndServe(opt.Addr, h)
}

// ParseFlags parses args into an Options struct.
func (m *Main) ParseFlags(args []string) (Options, error) {
	var opt Options
	fs := flag.NewFlagSet("", flag.ContinueOnError)
	fs.StringVar(&opt.Addr, "addr", ":1080", "bind address")
	if err := fs.Parse(args); err != nil {
		return opt, err
	}
	return opt, nil
}

// Options represents the command line options.
type Options struct {
	Addr string
}
