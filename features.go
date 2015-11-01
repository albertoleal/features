package features

import (
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

func (f *Features) IsActive(featureKey string) bool {
	ffk := engine.FeatureFlagKey{Key: featureKey}
	feature, err := f.ng.GetFeatureFlag(ffk)
	if err != nil {
		// TODO: add log here
		return false
	}
	return feature.Enabled == true
}

func (f *Features) IsInactive(featureKey string) bool {
	return !f.IsActive(featureKey)
}

func (f *Features) With(featureKey string, fn func()) {
	if f.IsActive(featureKey) {
		fn()
	}
}

func (f *Features) Without(featureKey string, fn func()) {
	if f.IsInactive(featureKey) {
		fn()
	}
}

// User has access if:
// - the feature is active;
// - the feature is inactive but the user has explicit access to it
func (f *Features) UserHasAccess(featureKey string, userId string) bool {
	ffk := engine.FeatureFlagKey{Key: featureKey}
	feature, err := f.ng.GetFeatureFlag(ffk)
	if err != nil {
		// TODO: add log here
		return false
	}

	if f.IsActive(featureKey) {
		return true
	}

	user := &engine.User{Id: userId}
	if feature.ContainsUser(user) {
		return true
	}

	return false
}
