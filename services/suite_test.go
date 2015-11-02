// Copyright 2015 Features authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package services_test

import (
	"testing"

	"github.com/albertoleal/features"
	"github.com/albertoleal/features/engine/memory"
	. "gopkg.in/check.v1"
)

type S struct {
	features *features.Features
}

var _ = Suite(&S{})

//Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

func (s *S) SetUpTest(c *C) {
	s.features = features.New(memory.New())
}
