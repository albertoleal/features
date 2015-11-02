// Copyright 2015 Features authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package services_test

import (
	"github.com/albertoleal/features/engine"
	"github.com/albertoleal/features/services"
	. "gopkg.in/check.v1"
)

func (s *S) TestCreateFeatureFlag(c *C) {
	feature := &engine.FeatureFlag{
		Key:     "Feature X",
		Enabled: false,
	}

	service := services.NewFeatureService(s.features)
	c.Check(service.CreateFeatureFlag(feature), IsNil)
}

func (s *S) TestCreateFeatureFlagWithInvalidData(c *C) {
	feature := &engine.FeatureFlag{
		Key:     "",
		Enabled: false,
	}

	service := services.NewFeatureService(s.features)
	c.Check(service.CreateFeatureFlag(feature), Not(IsNil))
}
