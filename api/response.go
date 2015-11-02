// Copyright 2015 Features authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package api

import (
	"encoding/json"
	"net/http"
)

type HTTPResponse struct {
	ContentType string      `json:"-"`
	StatusCode  int         `json:"-"`
	Body        interface{} `json:"body,omitempty"`
}

func (h *HTTPResponse) ToJson() []byte {
	h.ContentType = "application/json"

	r, err := json.Marshal(h.Body)
	if err != nil {
		return []byte(err.Error())
	}
	return r
}

func encodeResponse(rw http.ResponseWriter, response interface{}) error {
	resp := response.(HTTPResponse)
	var contentType string
	if contentType = resp.ContentType; contentType == "" {
		contentType = "application/json"
	}
	rw.Header().Set("Content-Type", contentType)

	rw.WriteHeader(resp.StatusCode)
	if resp.Body != nil {
		return json.NewEncoder(rw).Encode(resp.Body)
	}
	return nil
}

func handleErrorEncoder(rw http.ResponseWriter, err error) {
	switch err.(type) {
	case NotFoundError:
		erro := ErrorResponse{Type: E_NOT_FOUND, Description: err.Error()}
		NotFound(rw, erro)
	default:
		erro := ErrorResponse{Type: E_BAD_REQUEST, Description: err.Error()}
		BadRequest(rw, erro)
	}
}

func NotFound(rw http.ResponseWriter, body interface{}) {
	resp := HTTPResponse{StatusCode: http.StatusNotFound, Body: body}
	jsonResponse(rw, resp)
}

func BadRequest(rw http.ResponseWriter, body interface{}) {
	resp := HTTPResponse{StatusCode: http.StatusBadRequest, Body: body}
	jsonResponse(rw, resp)
}

func jsonResponse(rw http.ResponseWriter, resp HTTPResponse) {
	body := resp.ToJson()
	rw.Header().Set("Content-Type", resp.ContentType)
	rw.WriteHeader(resp.StatusCode)
	rw.Write([]byte(body))
}
