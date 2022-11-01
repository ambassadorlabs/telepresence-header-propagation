package main

import (
	"context"
	"errors"
	httptransport "github.com/go-kit/kit/transport/http"
	"net/url"
	"strings"
)

// StringService provides operations on strings.
type StringService interface {
	InitialUppercase(context.Context, string) (string, error)
	FinalUppercase(context.Context, string) (string, error)
	Count(context.Context, string) int
}

type stringService struct{}

func (stringService) InitialUppercase(ctx context.Context, s string) (string, error) {
	u, _ := url.Parse("http://localhost:8080/finaluppercase")
	client := httptransport.NewClient(
		"POST",
		u,
		encodeRequest,
		decodeUppercaseResponse,
		httptransport.ClientBefore(setHeaders),
	).Endpoint()

	resp, err := client(ctx, uppercaseRequest{S: s})
	if err != nil {
		return "", err
	}

	if resp.(uppercaseResponse).Err != "" {
		return "", errors.New(resp.(uppercaseResponse).Err)
	}
	return resp.(uppercaseResponse).V, nil
}

func (stringService) FinalUppercase(ctx context.Context, s string) (string, error) {
	if s == "" {
		return "", ErrEmpty
	}
	return strings.ToUpper(s), nil
}

func (stringService) Count(ctx context.Context, s string) int {
	return len(s)
}

// ErrEmpty is returned when an input string is empty.
var ErrEmpty = errors.New("empty string")

// ServiceMiddleware is a chainable behavior modifier for StringService.
type ServiceMiddleware func(StringService) StringService
