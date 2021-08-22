package vimego

import (
	"errors"
	"fmt"
)

var (
	ErrInvalidUrl     = errors.New("the URL is invalid")
	ErrDecodingFailed = errors.New("couldn't decode JSON")
	ErrParsingFailed  = errors.New("couldn't find config url")
)

type ErrUnexpectedStatusCode int

func (err ErrUnexpectedStatusCode) Error() string {
	return fmt.Sprintf("unexpected status code: %d", err)
}
