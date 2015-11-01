package engine

import (
	"errors"
)

type User struct {
	Id string
}

type FeatureFlag struct {
	Enabled bool    `json:"enabled"`
	Key     string  `json:"key"`
	Users   []*User `json:"users"`
}

func (ff *FeatureFlag) ContainsUser(user *User) bool {
	for _, u := range ff.Users {
		if u.Id == user.Id {
			return true
		}
	}
	return false
}

func NewFeatureFlag(key string, enabled bool, users []*User) (*FeatureFlag, error) {
	if key == "" {
		return nil, errors.New("Key cannot be empty.")
	}

	return &FeatureFlag{
		Key:     key,
		Enabled: enabled,
		Users:   users,
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
