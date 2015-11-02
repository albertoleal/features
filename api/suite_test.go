// Copyright 2015 Features authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package api_test

import (
	"net/http/httptest"
	"testing"

	"github.com/albertoleal/features/api"
	"github.com/albertoleal/features/engine"
	"github.com/albertoleal/features/engine/memory"
	"github.com/apihub/apihub/requests"
	. "gopkg.in/check.v1"
)

var httpClient requests.HTTPClient

func Test(t *testing.T) { TestingT(t) }

type S struct {
	api    *api.Api
	server *httptest.Server
	ng     engine.Engine
}

func (s *S) SetUpTest(c *C) {
	s.ng = memory.New()
	s.api = api.NewApi(s.ng)
	s.server = httptest.NewServer(s.api.Handler())
	httpClient = requests.NewHTTPClient(s.server.URL)
}

var _ = Suite(&S{})
