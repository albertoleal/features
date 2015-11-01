package engine

// Engine is an interface for storage.
type Engine interface {
	// GetFeatureFlags returns list of features registered in the storage engine.
	// Returns empty list otherwise.
	GetFeatureFlags() ([]FeatureFlag, error)
	// GetFeatureFlag returns feature by given key, or engine.NotFoundError when it's not found.
	GetFeatureFlag(FeatureFlagKey) (*FeatureFlag, error)
	// UpsertFeatureFlag updates or inserts the feature flag
	UpsertFeatureFlag(FeatureFlag) error
	// DeleteFeatureFlag deletes feature flag  by given key or returns engine.NotFoundError when it's not found.
	DeleteFeatureFlag(FeatureFlagKey) error
}
