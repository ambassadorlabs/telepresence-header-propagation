package main

import (
	"context"
	"net/http"
)

const telepresenceHeader string = "x-telepresence-id"

func extractHeaders(ctx context.Context, r *http.Request) context.Context {
	header := r.Header.Get(telepresenceHeader)
	ctx = context.WithValue(ctx, telepresenceHeader, header)
	return ctx
}

func setHeaders(ctx context.Context, r *http.Request) context.Context {
	r.Header.Set(telepresenceHeader, ctx.Value(telepresenceHeader).(string))
	return ctx
}
