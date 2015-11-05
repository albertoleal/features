// Copyright 2015 Features authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package api

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/albertoleal/features"
	"github.com/albertoleal/features/engine"
	"github.com/go-kit/kit/endpoint"
	"github.com/gorilla/mux"
	"golang.org/x/net/context"
)

type validationRequest struct {
	Key  string `json:"key"`
	User string `json:"user"`
}

func decodeValidationRequest(r *http.Request) (interface{}, error) {
	var request validationRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		switch {
		case err == io.EOF:
		case err != nil:
			return nil, err
		}
	}
	return request, nil
}

type featureFlagRequest struct {
	engine.FeatureFlag
}

func decodeFeatureFlagQueryString(r *http.Request) (interface{}, error) {
	return mux.Vars(r)["feature_key"], nil
}

func decodeFeatureFlagRequest(r *http.Request) (interface{}, error) {
	var request featureFlagRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		switch {
		case err == io.EOF:
		case err != nil:
			return nil, err
		}
	}
	return request, nil
}

func MakeCreate(feature features.Features) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(featureFlagRequest)
		ff := req.FeatureFlag

		if ff, _ := feature.Find(ff.Key); ff != nil {
			errResp := NewErrorResponse(E_BAD_REQUEST, "There's another feature for the same key value.")
			return HTTPResponse{StatusCode: http.StatusBadRequest, Body: errResp}, nil
		}

		err := feature.Save(ff)
		if err != nil {
			return nil, err
		}

		return HTTPResponse{StatusCode: http.StatusCreated, Body: ff}, nil
	}
}

func MakeUpdate(feature features.Features) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(featureFlagRequest)
		ff := req.FeatureFlag

		if ff, err := feature.Find(ff.Key); ff == nil {
			return nil, NewNotFoundError(err)
		}

		err := feature.Save(ff)
		if err != nil {
			return nil, err
		}

		return HTTPResponse{StatusCode: http.StatusOK, Body: ff}, nil
	}
}

func MakeDelete(feature features.Features) endpoint.Endpoint {
	return func(ctx context.Context, feature_key interface{}) (interface{}, error) {
		fk := feature_key.(string)
		err := feature.Delete(fk)
		if err != nil {
			return nil, err
		}

		return HTTPResponse{StatusCode: http.StatusNoContent}, nil
	}
}

func MakeFind(feature features.Features) endpoint.Endpoint {
	return func(ctx context.Context, feature_key interface{}) (interface{}, error) {
		fk := feature_key.(string)
		ff, err := feature.Find(fk)
		if err != nil {
			return nil, err
		}

		return HTTPResponse{StatusCode: http.StatusOK, Body: ff}, nil
	}
}

func MakeList(feature features.Features) endpoint.Endpoint {
	return func(ctx context.Context, feature_key interface{}) (interface{}, error) {

		ffs, err := feature.List()
		if err != nil {
			return nil, err
		}

		cs := &CollectionSerializer{
			Items: ffs,
			Count: len(ffs),
		}
		return HTTPResponse{StatusCode: http.StatusOK, Body: cs}, nil
	}
}

func MakeValidate(feature features.Features) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(validationRequest)

		ff, err := feature.Find(req.Key)
		if err != nil {
			return HTTPResponse{StatusCode: http.StatusForbidden}, nil
		}

		var access bool
		if ff.Percentage > 0 || len(ff.Users) > 0 {
			if access = feature.UserHasAccess(req.Key, req.User); access {
				return HTTPResponse{StatusCode: http.StatusOK}, nil
			} else {
				return HTTPResponse{StatusCode: http.StatusForbidden}, nil
			}
		}

		if access, _ = feature.IsEnabled(req.Key); access {
			return HTTPResponse{StatusCode: http.StatusOK}, nil
		}

		return HTTPResponse{StatusCode: http.StatusForbidden}, nil
	}
}
