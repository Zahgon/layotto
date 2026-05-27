package exporter_iml

import (
	"mosn.io/layotto/components/trace"
)

func init() {
	trace.RegisterExporter("stdout", &StdoutExporter{})
}

// StdoutExporter is the implementation of Exporter, export span information to log
type StdoutExporter struct{}

var _ trace.Exporter = &StdoutExporter{}

const msg = "%s, AppName: %+s, Method: %s, TraceId: %s, SpanId: %s, ParentSpanId:%s, Time: [%s ->  %s], processTime: %+v, result: %+v, extraInfo:[ %+v ]"

// ExportSpan implements the open census exporter interface.
func (e *StdoutExporter) ExportSpan(sd *trace.Span) { _ = "STUB: not implemented"; return }
