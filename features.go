// Copyright 2015 Features authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package features

import (
	"errors"

	"github.com/albertoleal/features/engine"
)

// type FeatureArgs struct {
// 	Key  engine.FeatureFlagKey
// 	User engine.User
// }

type Features struct {
	ng engine.Engine
}

func New(ng engine.Engine) *Features {
	return &Features{ng: ng}
}

func (f *Features) Save(feature engine.FeatureFlag) error {
	return f.ng.UpsertFeatureFlag(feature)
}

func (f *Features) Delete(featureKey string) error {
	ffk := engine.FeatureFlagKey{Key: featureKey}
	return f.ng.DeleteFeatureFlag(ffk)
}

func (f *Features) IsEnabled(featureKey string) (bool, error) {
	ffk := engine.FeatureFlagKey{Key: featureKey}
	feature, err := f.ng.GetFeatureFlag(ffk)
	if err != nil {
		// TODO: add log here
		return false, err
	}

	if feature.Percentage > 0 {
		// TODO: add log here
		return false, errors.New("Percentage is defined. Call `.UserHasAccess` instead.")
	}

	if len(feature.Users) > 0 {
		// TODO: add log here
		return false, errors.New("Users not empty. Call `.UserHasAccess` instead.")
	}

	return feature.Enabled == true, nil
}

func (f *Features) IsDisabled(featureKey string) (bool, error) {
	out, err := f.IsEnabled(featureKey)
	return !out, err
}

func (f *Features) With(featureKey string, fn func()) {
	if ok, err := f.IsEnabled(featureKey); ok && err == nil {
		fn()
	}
}

func (f *Features) Without(featureKey string, fn func()) {
	if ok, _ := f.IsDisabled(featureKey); ok {
		fn()
	}
}

// User has access if:
// - the feature is active;
// - the feature is inactive but the user has explicit access to it;
// - the feature is active for a percentage of users.
func (f *Features) UserHasAccess(featureKey string, userId string) bool {
	ffk := engine.FeatureFlagKey{Key: featureKey}
	feature, err := f.ng.GetFeatureFlag(ffk)
	if err != nil {
		// TODO: add log here
		return false
	}

	// Active
	if ok, _ := f.IsEnabled(featureKey); ok {
		return true
	}

	// Specific users
	user := &engine.User{Id: userId}
	if feature.ContainsUser(user) {
		return true
	}

	// Percentage of users
	if feature.UserInPercentage(user) {
		return feature.Enabled
	}

	return false
}

func (f *Features) Valid(ff *engine.FeatureFlag) (bool, error) {
	ffk := engine.FeatureFlagKey{Key: ff.Key}
	if feature, _ := f.ng.GetFeatureFlag(ffk); feature != nil {
		return false, errors.New("There's another feature for the same key value.")
	}

	if ff.Key == "" {
		return false, errors.New("Key cannot be empty.")
	}

	return true, nil
}
