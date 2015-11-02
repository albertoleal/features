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

func (f *Features) IsActive(featureKey string) (bool, error) {
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

	return feature.Enabled == true, nil
}

func (f *Features) IsInactive(featureKey string) (bool, error) {
	out, err := f.IsActive(featureKey)
	return !out, err
}

func (f *Features) With(featureKey string, fn func()) {
	if ok, _ := f.IsActive(featureKey); ok {
		fn()
	}
}

func (f *Features) Without(featureKey string, fn func()) {
	if ok, _ := f.IsInactive(featureKey); ok {
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
	if ok, _ := f.IsActive(featureKey); ok {
		return true
	}

	// Specific users
	user := &engine.User{Id: userId}
	if feature.ContainsUser(user) {
		return true
	}

	// Percentage of users
	if feature.UserInPercentage(user) {
		return true
	}

	return false
}
