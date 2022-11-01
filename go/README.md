# Header Forwarding Example - Go

## Credit where it's due

This example is built using [go-kit](https://github.com/go-kit/kit), a popular enterprise microservice framework for Go,
and heavily leans on the [stringsvc example](https://github.com/go-kit/examples/tree/master/stringsvc2) provided by
that project.

## How the example works

The service has an endpoint `/uppercase` that uppercases a string sent to it via a POST request with a JSON body.
Internally, the service makes a second HTTP call to `/finaluppercase` which does the actual uppercasing and sends the
new string back up the chain. To test it, you can run

```shell
$ go run *.go
```

and then make this request:

```shell
curl --request POST 'http://localhost:8080/uppercase' \
--header 'x-telepresence-id: 1234567890' \
--header 'Content-Type: application/json' \
--data-raw '{
    "s": "hello, world"
}'
```

In the console output of the running process you'll see

```
listen=:8080 caller=logging.go:35 method=finalUppercase telepresenceHeader=1234567890
listen=:8080 caller=logging.go:23 method=uppercase telepresenceHeader=1234567890
```

This output shows that the `x-telepresence-id` header has been propagated from your call to `/uppercase`, to the second
endpoint `/finaluppercase`.

### How it's done: context.Context

The most effective way to propagate headers in Go is through a `context`. `go-kit` provides some convenience methods for
manipulating incoming and outgoing requests. At line 30 of `main.go`, we add `httptransport.ServerBefore(extractHeaders)`
to the stack of function calls on an incoming request. `extractHeaders` is in `context_headers.go`:

```go
func extractHeaders(ctx context.Context, r *http.Request) context.Context {
	header := r.Header.Get("X-Telepresence-Id")
	ctx = context.WithValue(ctx, "x-telepresence-id", header)
	return ctx
}
```

`ServerBefore` calls any functions passed to it before the request is passed further down the stack. So we get the header
we want to propagate and add it to the context that is passed down through the call stack. In this case we are only grabbing
the `X-Telepresence-Id` header but we could add any others that we care about, or even all of them.

To make the request to `/finaluppercase`, in `service.go`, in the `InitialUppercase` function, we create a new HTTP client:

```go
client := httptransport.NewClient(
		"POST",
		u,
		encodeRequest,
		decodeUppercaseResponse,
		httptransport.ClientBefore(setHeaders),
	).Endpoint()
```

The `ClientBefore` convenience method lets us pass functions that run before a request is sent. The `setHeaders` method
we are passing looks like

```go
func setHeaders(ctx context.Context, r *http.Request) context.Context {
	r.Header.Set("x-telepresence-id", ctx.Value("x-telepresence-id").(string))
	return ctx
}
```

Similar to extracting the headers, we're setting the oen we want but could do any or all headers we want to propagate.

## Summary

Regardless of how your Go services are structured, the key to propagating headers is to retrieve them from the request
as soon as possible, and attach them to a context that is passed down through your call stack. Then you can add them to
any outgoing requests to upstream services.
