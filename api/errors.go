// Copyright 2015 Features authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package errors contains common business errors to be user all over the server.
package api

import "errors"

const (
	E_BAD_REQUEST string = "bad_request"
	E_NOT_FOUND   string = "not_found"
)

var (
	ErrNotFound = errors.New("The resource requested does not exist.")
)

type ErrorResponse struct {
	Type        string `json:"error,omitempty"`
	Description string `json:"error_description,omitempty"`
}

func (err ErrorResponse) Error() string {
	return err.Description
}

func NewErrorResponse(errType, description string) ErrorResponse {
	return ErrorResponse{Type: errType, Description: description}
}

type NotFoundError struct {
	description error
}

func NewNotFoundError(err error) NotFoundError {
	return NotFoundError{description: err}
}

func (err NotFoundError) Error() string {
	return err.description.Error()
}
