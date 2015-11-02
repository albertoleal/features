// Copyright 2015 Features authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package services

import (
	"encoding/json"
)

type ErrorService interface {
	ToJson() ([]byte, error)
}

type errorService struct {
	Code        string `json:"code,omitempty"`
	Description string `json:"error_description,omitempty"`
	Type        string `json:"error,omitempty"`
}

func (e errorService) ToJson() ([]byte, error) {
	json, err := json.Marshal(e)
	if err != nil {
		return []byte(err.Error()), err
	}

	return json, nil
}
