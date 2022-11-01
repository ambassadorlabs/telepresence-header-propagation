package main

import (
	"context"
	"net/http"
)

func extractHeaders(ctx context.Context, r *http.Request) context.Context {
	header := r.Header.Get("X-Telepresence-Id")
	ctx = context.WithValue(ctx, "x-telepresence-id", header)
	return ctx
}

func setHeaders(ctx context.Context, r *http.Request) context.Context {
	r.Header.Set("x-telepresence-id", ctx.Value("x-telepresence-id").(string))
	return ctx
}
