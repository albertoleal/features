// Copyright 2015 Features authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package services

import (
	"github.com/albertoleal/features"
	"github.com/albertoleal/features/engine"
)

type FeatureService interface {
	CreateFeatureFlag(*engine.FeatureFlag) error
}

type featureService struct {
	features *features.Features
}

func NewFeatureService(ff *features.Features) FeatureService {
	return &featureService{features: ff}
}

// `CreateFeatureFlag` creates a new feature flag.
//
// It requires to inform the following field(s): Key.
// It is not allowed to create two feature flags with the same key.
// It returns an error if when it fails.
func (service featureService) CreateFeatureFlag(ff *engine.FeatureFlag) error {
	if _, err := service.features.Valid(ff); err != nil {
		return err
	}

	return service.features.Save(*ff)
}
