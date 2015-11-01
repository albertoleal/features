package engine

import (
	"errors"
)

type FeatureFlag struct {
	Enabled bool   `json:"enabled"`
	Key     string `json:"key"`
}

func NewFeatureFlag(key string, enabled bool) (*FeatureFlag, error) {
	if key == "" {
		return nil, errors.New("Key cannot be empty.")
	}

	return &FeatureFlag{
		Key:     key,
		Enabled: enabled,
	}, nil
}

type FeatureFlagKey struct {
	Key string
}

type NotFoundError struct {
	Message string
}

func (n *NotFoundError) Error() string {
	if n.Message != "" {
		return n.Message
	}

	return "Feature Flag not found."
}
