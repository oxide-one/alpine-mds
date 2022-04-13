package main

import (
	"errors"
	"fmt"
	"time"
)

type Request struct {
	Location string `json:"location"`
}

type Response struct {
	StatusCode int               `json:"statusCode,omitempty"`
	Headers    map[string]string `json:"headers,omitempty"`
	Body       string            `json:"body,omitempty"`
}

func Main(in Request) (*Response, error) {
	if in.Location == "" {
		return nil, errors.New("location must be passed")
	}

	return &Response{
		Body: fmt.Sprintf("Hello, from %s, it is now %s", in.Location, time.Now().String()),
	}, nil
}
