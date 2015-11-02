// Copyright 2015 Features authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package api_test

import (
	"net/http"

	"github.com/albertoleal/features/api"
	"github.com/albertoleal/features/engine"
	. "gopkg.in/check.v1"
)

func (s *S) TestToJsonWithError(c *C) {
	erro := api.ErrorResponse{
		Type:        "invalid_request",
		Description: "The request is missing a required parameter.",
	}

	err := &api.HTTPResponse{
		StatusCode: http.StatusBadRequest,
		Body:       erro,
	}
	c.Assert(string(err.ToJson()), Equals, `{"error":"invalid_request","error_description":"The request is missing a required parameter."}`)
}

func (s *S) TestToJsonWithUser(c *C) {
	ff := &engine.FeatureFlag{Key: "login_via_email", Enabled: true, Percentage: 85}
	err := &api.HTTPResponse{
		StatusCode: http.StatusOK,
		Body:       ff,
	}
	c.Assert(string(err.ToJson()), Equals, `{"enabled":true,"key":"login_via_email","percentage":85}`)
}
