// Package gogsi is a game state integration (GSI) library for the Dota 2 client. More information is available
// on github: https://github.com/mammothbane/gogsi/blob/master/README.md
package gogsi

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"
)

// (.+?)://   - scheme (ignored)
// ([^/:]*?)  - host is everything up to the first colon or slash
// (:[0-9]+)  - port comes after a colon and is only numerals
// (\/.*)     - everything after the slash (including slash) is the path
var uriRegex = regexp.MustCompile(`^(?:(.+?)://)?([^/:]*)(:[0-9]+)?(\/.*)?$`)

// Listen is the main entrypoint into gogsi.
//
// Url
//
// Can be any portion of a URL.
// Each of the following is valid and  equivalent to all the others:
//  gogsi.Listen("http://localhost:3000", handler)
//  gogsi.Listen("localhost:3000", handler)
//  gogsi.Listen(":3000", handler)
//  gogsi.Listen("", handler)
// A path can optionally follow:
//  gogsi.Listen(":3000/my/custom/url")
//
// Defaults
//
// If not present, the host is assumed to be 'localhost', the path is assumed to be "/",
// and the port is assumed to be 3000. The scheme is ignored and not validated&mdash;
// "my_scheme://localhost:3000" is equivalent to "http://localhost:3000".
//
// Closure
//
// fn contains the logic for whatever you want to do on each update. It's called as part of
// an HTTP handler, so a goroutine is spawned for each invocation. As such, heavy work within fn
// is not recommended, at least unless GSi is configured with a high buffer value. Dota will wait
// to receive a 2XX response until it times out (and it will retry if it doesn't receive a 2XX), so
// processing for too long could quickly deadlock your system or run you out of memory.
// If your closure returns an error, the server will respond to the Dota client with a 500, triggering
// a retry. A nil return from your closure always results in a 200 OK.
func Listen(url string, fn func(state *State) error) error {
	matches := uriRegex.FindStringSubmatch(url)

	host := matches[2]
	if len(matches[2]) == 0 {
		host = "localhost"
	}

	port := matches[3]
	if len(matches[3]) == 0 {
		port = ":3000"
	}

	path := matches[4]
	if len(matches[4]) == 0 {
		path = "/"
	}

	http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		gw := gogsiWriter{w}

		if r.Method != "POST" {
			gw.sendErr(http.StatusBadRequest, "bad request method: %v", r.Method)
			return
		}

		if r.Header.Get("Content-Type") != "application/json" {
			gw.sendErr(http.StatusBadRequest, "bad content type for request: got '%v', required 'application/json'", r.Header.Get("Content-Type"))
			return
		}

		var state State
		if err := json.NewDecoder(r.Body).Decode(&state); err != nil {
			gw.sendErr(http.StatusNoContent, "unable to decode request body: %v", err)
			return
		}

		if err := fn(&state); err != nil {
			gw.sendErr(http.StatusInternalServerError, "responding to game update: %v", err)
			return
		}
		w.WriteHeader(http.StatusOK)
	})
	fmt.Println("listening on", host+port, "path", path)
	return http.ListenAndServe(host+port, nil)
}

type gogsiWriter struct {
	http.ResponseWriter
}

func (g gogsiWriter) sendErr(code int, message string, args ...interface{}) {
	log.Printf("ERROR: "+message, args...)
	http.Error(g, fmt.Sprintf(message, args...), code)
}
