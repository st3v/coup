package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/st3v/coup/service"
)

var addr string

func init() {
	log.SetFlags(0)
	flag.StringVar(&addr, "listen", "", "Address to listen on for HTTP POST requests.")
}

func main() {
	flag.Parse()

	if addr != "" {
		http.Handle("/", handleHTTP())
		fmt.Printf("Listening on %s...\n", addr)
		log.Fatal(http.ListenAndServe(addr, nil))
	}

	if err := handle(os.Stdout, os.Stdin); err != nil {
		log.Fatal(err)
	}
}

// handleHTTP returns an http.Handler that is a simple wrapper around the
// the more generic handle function.
func handleHTTP() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if req.Method != "POST" {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
			return
		}

		defer req.Body.Close()
		if err := handle(w, req.Body); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	})
}

// handle reads request from io.Reader r, handles the request, and writes the
// response to io.Writer w. Returns an error if any of the 3 steps fails.
func handle(w io.Writer, r io.Reader) error {
	var req service.Request
	if err := json.NewDecoder(r).Decode(&req); err != nil {
		return fmt.Errorf("error decoding payload: %v", err)
	}

	resp, err := req.Handle()
	if err != nil {
		return fmt.Errorf("error handling request: %v", err)
	}

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		return fmt.Errorf("error encoding response: %v", err)
	}

	return nil
}
