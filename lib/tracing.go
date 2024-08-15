package lib

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/opentracing/opentracing-go/ext"
	"io"
	"net/http"
	"runtime"
	"strings"
	"time"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
	"github.com/uber/jaeger-client-go/config"
)

// InitJaeger initializes and returns a Jaeger Tracer with 100% sampling and logging all spans to stdout.
func InitJaeger(service string) (opentracing.Tracer, io.Closer, error) {
	cfg, err := getDefaultConfig(service).FromEnv()
	if err != nil {
		return nil, nil, fmt.Errorf("could not parse Jaeger env vars: %w", err)
	}

	tracer, closer, err := cfg.NewTracer()
	if err != nil {
		return nil, nil, fmt.Errorf("could not initialize Jaeger tracer: %w", err)
	}

	return tracer, closer, nil
}

func getDefaultConfig(service string) *config.Configuration {
	return &config.Configuration{
		ServiceName: service,
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans:            true,
			BufferFlushInterval: 1 * time.Second,
		},
	}
}

// CreateRootSpan creates a new root span from an incoming HTTP request and enriches it with additional tags and logs.
func CreateRootSpan(r *http.Request, operationName string) opentracing.Span {
	spanCtx, _ := opentracing.GlobalTracer().Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(r.Header))
	span := opentracing.StartSpan(operationName, ext.RPCServerOption(spanCtx))
	span.SetTag("http.method", r.Method)
	span.SetTag("http.url", r.URL.String())

	// Add user agent
	userAgent := r.UserAgent()
	span.SetTag("http.user_agent", userAgent)

	// Add request headers
	for key, values := range r.Header {
		span.SetTag("http.header."+key, strings.Join(values, ", "))
	}

	addCallerTag(span)
	return span
}

// CreateChildSpan creates a new OpenTracing span as a child of the current span.
func CreateChildSpan(ctx context.Context, name string) opentracing.Span {
	parentSpan := opentracing.SpanFromContext(ctx)
	sp := startSpan(ctx, parentSpan, name)
	return sp
}

// CreateSubChildSpan creates a new OpenTracing span as a child of the provided parent span.
func CreateSubChildSpan(parentSpan opentracing.Span, name string) opentracing.Span {
	sp := startSpan(context.Background(), parentSpan, name)
	return sp
}

// LogRequest logs a request object to the span.
func LogRequest(sp opentracing.Span, req interface{}) {
	sp.LogFields(log.Object("request", stringify(req)))
}

// LogObject logs an object to the span.
func LogObject(sp opentracing.Span, name string, obj interface{}) {
	sp.LogFields(log.Object(name, stringify(obj)))
}

// LogResponse logs a response object to the span.
func LogResponse(sp opentracing.Span, resp interface{}) {
	sp.LogFields(log.Object("response", stringify(resp)))
}

// LogError logs an error to the span.
func LogError(sp opentracing.Span, err error) {
	sp.LogFields(log.Object("error", err))
	sp.SetTag("error", true)
}

// Helper functions
func startSpan(ctx context.Context, parentSpan opentracing.Span, name string) opentracing.Span {
	sp := opentracing.StartSpan(name, opentracing.ChildOf(parentSpan.Context()))
	sp.SetTag("name", name)
	addCallerTag(sp)
	return sp
}

func addCallerTag(sp opentracing.Span) {
	pc := make([]uintptr, 15)
	n := runtime.Callers(3, pc)
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()
	callerDetails := fmt.Sprintf("%s - %s#%d", frame.Function, frame.File, frame.Line)
	sp.SetTag("caller", callerDetails)
}

func stringify(data interface{}) string {
	dataByte, _ := json.Marshal(data)
	return string(dataByte)
}
