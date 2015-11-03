// Copyright 2015 Features authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package features

import (
	"errors"

	"github.com/albertoleal/features/engine"
	utils "github.com/mrvdot/golang-utils"
)

type Features interface {
	Save(engine.FeatureFlag) error
	List() ([]engine.FeatureFlag, error)
	Find(featureKey string) (*engine.FeatureFlag, error)
	Delete(featureKey string) error
	IsEnabled(featureKey string) (bool, error)
	IsDisabled(featureKey string) (bool, error)
	With(featureKey string, fn func())
	Without(featureKey string, fn func())
	UserHasAccess(featureKey string, userId string) bool
}

type features struct {
	ng engine.Engine
}

func New(ng engine.Engine) Features {
	return &features{ng: ng}
}

// `Save` creates a new feature flag.
//
// It requires to inform the following field(s): Key.
// It returns an error if when it fails.
func (f *features) Save(feature engine.FeatureFlag) error {
	feature.Key = utils.GenerateSlug(feature.Key)
	if _, err := f.valid(&feature); err != nil {
		return err
	}

	return f.ng.UpsertFeatureFlag(feature)
}

func (f *features) Delete(featureKey string) error {
	ffk := engine.FeatureFlagKey{Key: featureKey}
	return f.ng.DeleteFeatureFlag(ffk)
}

func (f *features) Find(featureKey string) (*engine.FeatureFlag, error) {
	ffk := engine.FeatureFlagKey{Key: featureKey}
	return f.ng.GetFeatureFlag(ffk)
}

func (f *features) List() ([]engine.FeatureFlag, error) {
	return f.ng.GetFeatureFlags()
}

func (f *features) IsEnabled(featureKey string) (bool, error) {
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

func (f *features) IsDisabled(featureKey string) (bool, error) {
	out, err := f.IsEnabled(featureKey)
	return !out, err
}

func (f *features) With(featureKey string, fn func()) {
	if ok, err := f.IsEnabled(featureKey); ok && err == nil {
		fn()
	}
}

func (f *features) Without(featureKey string, fn func()) {
	if ok, _ := f.IsDisabled(featureKey); ok {
		fn()
	}
}

// User has access if:
// - the feature is active;
// - the feature is inactive but the user has explicit access to it;
// - the feature is active for a percentage of users.
func (f *features) UserHasAccess(featureKey string, userId string) bool {
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

func (f *features) valid(ff *engine.FeatureFlag) (bool, error) {
	if ff.Key == "" {
		return false, errors.New("Key cannot be empty.")
	}

	return true, nil
}
