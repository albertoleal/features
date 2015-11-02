// Copyright 2015 Features authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package engine

import (
	"errors"
	"hash/crc32"
)

type User struct {
	Id string
}

type FeatureFlag struct {
	Enabled    bool    `json:"enabled"`
	Key        string  `json:"key"`
	Users      []*User `json:"users,omitempty"`
	Percentage uint32  `json:"percentage"`
}

func (ff *FeatureFlag) ContainsUser(user *User) bool {
	for _, u := range ff.Users {
		if u.Id == user.Id {
			return true
		}
	}
	return false
}

func (ff *FeatureFlag) UserInPercentage(user *User) bool {
	return crc32.ChecksumIEEE([]byte(user.Id))%100 <= ff.Percentage
}

func NewFeatureFlag(key string, enabled bool, users []*User, percentage uint32) (*FeatureFlag, error) {
	if key == "" {
		return nil, errors.New("Key cannot be empty.")
	}

	return &FeatureFlag{
		Key:        key,
		Enabled:    enabled,
		Users:      users,
		Percentage: percentage,
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

	return "Feature flag not found."
}
