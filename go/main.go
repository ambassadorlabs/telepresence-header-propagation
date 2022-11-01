package main

import (
	"flag"
	"net/http"
	"os"

	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
)

func main() {
	var (
		listen = flag.String("listen", ":8080", "HTTP listen address")
	)
	flag.Parse()

	var logger log.Logger
	logger = log.NewLogfmtLogger(os.Stderr)
	logger = log.With(logger, "listen", *listen, "caller", log.DefaultCaller)

	var svc StringService
	svc = stringService{}
	svc = loggingMiddleware(logger)(svc)

	uppercaseHandler := httptransport.NewServer(
		makeUppercaseEndpoint(svc),
		decodeUppercaseRequest,
		encodeResponse,
		httptransport.ServerBefore(extractHeaders),
	)
	finalUppercaseHandler := httptransport.NewServer(
		makeFinalUppercaseEndpoint(svc),
		decodeUppercaseRequest,
		encodeResponse,
		httptransport.ServerBefore(extractHeaders),
	)
	countHandler := httptransport.NewServer(
		makeCountEndpoint(svc),
		decodeCountRequest,
		encodeResponse,
		httptransport.ServerBefore(extractHeaders),
	)

	http.Handle("/uppercase", uppercaseHandler)
	http.Handle("/finaluppercase", finalUppercaseHandler)
	http.Handle("/count", countHandler)
	logger.Log("msg", "HTTP", "addr", *listen)
	logger.Log("err", http.ListenAndServe(*listen, nil))
}
