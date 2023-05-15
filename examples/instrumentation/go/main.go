package main

import (
	"context"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"

	// these otel packages need to be imported
	"go.opentelemetry.io/contrib/instrumentation/github.com/gorilla/mux/otelmux"
	"go.opentelemetry.io/otel"
	stdout "go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/propagation"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

// the following code establishes an otel tracer and a function to initialize it
var tracer = otel.Tracer("uppercase-server")

func initTracer() (*sdktrace.TracerProvider, error) {
	exporter, err := stdout.New(stdout.WithPrettyPrint())
	if err != nil {
		return nil, err
	}
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithBatcher(exporter),
		)
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))
	return tp, nil
}

func main() {
	// we initialize the tracer
	tp, err := initTracer()
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := tp.Shutdown(context.Background()); err != nil {
			log.Printf("Error shutting down tracer provider: %v", err)
		}
	}()

	r := mux.NewRouter()
	
	// and use it as middleware for Gorilla mux, which will cause it to
	// add all otel info (including incoming Baggage) onto the context
	// for use in further function calls
	r.Use(otelmux.Middleware("uppercase"))

	r.HandleFunc("/finalupper", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		subject := r.URL.Query().Get("subject")
		uppercased := strings.ToUpper(subject)
		upperBytes := []byte(uppercased)
		w.WriteHeader(200)
		w.Write(upperBytes)
	}))
	
	http.Handle("/", r)
	_ = http.ListenAndServe(":8080", nil)
}
