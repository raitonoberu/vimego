package vimego

import (
	"errors"
	"fmt"
)

var (
	ErrInvalidUrl    = errors.New("the URL is invalid")
	ErrParsingFailed = errors.New("couldn't get config")
)

type ErrUnexpectedStatusCode int

func (err ErrUnexpectedStatusCode) Error() string {
	return fmt.Sprintf("unexpected status code: %d", err)
}
