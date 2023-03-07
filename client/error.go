// Copyright 2023 CodeMaker AI Inc. All rights reserved.

package client

type ClientError interface {
	error
	Unwrap() error
}

type clientError struct {
	ClientError
	message string
	cause   error
}

func NewClientError(message string) ClientError {
	return &clientError{
		message: message,
	}
}

func NewClientErrorWithCause(message string, cause error) ClientError {
	return &clientError{
		message: message,
		cause:   cause,
	}
}

func (e *clientError) Error() string {
	return e.message
}

func (e *clientError) Unwrap() error {
	return e.cause
}
