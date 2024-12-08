package trace

import (
	"context"
	"io"
	"net/http"
	_ "unsafe"

	traceinner "golang.org/x/net/trace/trace"
)

var (
	//go:linkname DebugUseAfterFinish golang.org/x/net/trace/trace.DebugUseAfterFinish
	DebugUseAfterFinish bool

	//go:linkname AuthRequest golang.org/x/net/trace/trace.AuthRequest
	AuthRequest func(req *http.Request) (any bool, sensitive bool)
)

type (
	Trace    = traceinner.Trace
	EventLog = traceinner.EventLog
)

func New(family, title string) Trace {
	return traceinner.New(family, title)
}

func NewEventLog(family, title string) EventLog {
	return traceinner.NewEventLog(family, title)
}

func NewContext(ctx context.Context, tr Trace) context.Context {
	return traceinner.NewContext(ctx, tr)
}

func FromContext(ctx context.Context) (tr Trace, ok bool) {
	return traceinner.FromContext(ctx)
}

func Events(w http.ResponseWriter, req *http.Request) {
	traceinner.Events(w, req)
}

func Traces(w http.ResponseWriter, req *http.Request) {
	traceinner.Traces(w, req)
}

func Render(w io.Writer, req *http.Request, sensitive bool) {
	traceinner.Render(w, req, sensitive)
}

func RenderEvents(w http.ResponseWriter, req *http.Request, sensitive bool) {
	traceinner.RenderEvents(w, req, sensitive)
}
