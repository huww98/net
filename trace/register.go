package trace

import (
	"net/http"
	"net/url"

	traceinner "golang.org/x/net/trace/trace"
)

// HTTP ServeMux paths.
const (
	debugRequestsPath = "/debug/requests"
	debugEventsPath   = "/debug/events"
)

func init() {
	_, pat := http.DefaultServeMux.Handler(&http.Request{URL: &url.URL{Path: debugRequestsPath}})
	if pat == debugRequestsPath {
		panic("/debug/requests is already registered. You may have two independent copies of " +
			"golang.org/x/net/trace in your binary, trying to maintain separate state. This may " +
			"involve a vendored copy of golang.org/x/net/trace.")
	}

	// TODO(jbd): Serve Traces from /debug/traces in the future?
	// There is no requirement for a request to be present to have traces.
	http.HandleFunc(debugRequestsPath, traceinner.Traces)
	http.HandleFunc(debugEventsPath, traceinner.Events)
}
