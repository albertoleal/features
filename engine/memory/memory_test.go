// Copyright 2015 Features authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package memory

import (
	"testing"

	"github.com/albertoleal/features/engine/test"
	. "gopkg.in/check.v1"
)

func TestMemory(t *testing.T) {
	Suite(&test.EngineSuite{Engine: New()})
	TestingT(t)
}
