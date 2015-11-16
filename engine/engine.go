// Copyright 2015 Features authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package engine

// Storage interface.
type Engine interface {
	// GetFeatureFlags returns list of features or an empty list otherwise.
	GetFeatureFlags() ([]FeatureFlag, error)
	// GetFeatureFlag returns an specific feature or engine.NotFoundError when it's not found.
	GetFeatureFlag(FeatureFlagKey) (*FeatureFlag, error)
	// UpsertFeatureFlag updates or inserts the feature flag.
	UpsertFeatureFlag(FeatureFlag) error
	// DeleteFeatureFlag deletes an specific feature flag or returns engine.NotFoundError when it's not found.
	DeleteFeatureFlag(FeatureFlagKey) error
}
