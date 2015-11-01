package features

import (
	"github.com/albertoleal/features/engine"
)

type Features struct {
	ng engine.Engine
}

func New(ng engine.Engine) *Features {
	return &Features{ng: ng}
}

func (f *Features) AddFeature(feature engine.FeatureFlag) error {
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
