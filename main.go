package main

import (
	"log"
	"net"
	"net/http"
	"os"
	"strings"
)

func main() {

	bind := "127.0.0.1:8080"
	if len(os.Args) >= 2 {
		bind = os.Args[1]
	}

	log.Printf("listening on %s", bind)

	http.ListenAndServe(bind, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" && r.URL.Path != "" {
			http.NotFound(w, r)
			return
		}

		w.Write([]byte(GetRemoteIP(r, true)))
	}))
}

func GetRemoteIP(r *http.Request, trustXForwardedFor bool) (source string) {
	// if we trust XForwardedFor and have such a header
	if xff := r.Header.Get("X-Forwarded-For"); xff != "" && trustXForwardedFor {
		source, _, _ = strings.Cut(xff, ",")
		return source
	}

	source = r.RemoteAddr

	// try to trim off a trailing port
	if host, _, err := net.SplitHostPort(source); err == nil {
		source = host
	}
	return
}
