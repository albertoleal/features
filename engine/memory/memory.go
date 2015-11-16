// Copyright 2015 Features authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package memory

import (
	"sync"

	"github.com/albertoleal/features/engine"
)

type Memory struct {
	FeatureFlags map[engine.FeatureFlagKey]engine.FeatureFlag
	mtx          sync.RWMutex
}

func New() engine.Engine {
	return &Memory{
		FeatureFlags: make(map[engine.FeatureFlagKey]engine.FeatureFlag),
	}
}

func (m *Memory) GetFeatureFlags() ([]engine.FeatureFlag, error) {
	ffs := make([]engine.FeatureFlag, 0, len(m.FeatureFlags))
	for _, ff := range m.FeatureFlags {
		ffs = append(ffs, ff)
	}
	return ffs, nil
}

func (m *Memory) GetFeatureFlag(k engine.FeatureFlagKey) (*engine.FeatureFlag, error) {
	m.mtx.RLock()
	defer m.mtx.RUnlock()
	ff, ok := m.FeatureFlags[k]
	if !ok {
		return nil, &engine.NotFoundError{}
	}
	return &ff, nil
}

func (m *Memory) UpsertFeatureFlag(ff engine.FeatureFlag) error {
	m.mtx.Lock()
	m.FeatureFlags[engine.FeatureFlagKey{Key: ff.Key}] = ff
	m.mtx.Unlock()
	return nil
}

func (m *Memory) DeleteFeatureFlag(k engine.FeatureFlagKey) error {
	if _, ok := m.FeatureFlags[k]; !ok {
		return &engine.NotFoundError{}
	}
	m.mtx.Lock()
	delete(m.FeatureFlags, k)
	m.mtx.Unlock()
	return nil
}
