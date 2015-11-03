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

func makeCreateFeatureFlag(feature features.Features) endpoint.Endpoint {
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

func makeUpdateFeatureFlag(feature features.Features) endpoint.Endpoint {
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

func makeDeleteFeatureFlag(feature features.Features) endpoint.Endpoint {
	return func(ctx context.Context, feature_key interface{}) (interface{}, error) {
		fk := feature_key.(string)
		err := feature.Delete(fk)
		if err != nil {
			return nil, err
		}

		return HTTPResponse{StatusCode: http.StatusNoContent}, nil
	}
}

func makeFindFeatureFlag(feature features.Features) endpoint.Endpoint {
	return func(ctx context.Context, feature_key interface{}) (interface{}, error) {
		fk := feature_key.(string)
		ff, err := feature.Find(fk)
		if err != nil {
			return nil, err
		}

		return HTTPResponse{StatusCode: http.StatusOK, Body: ff}, nil
	}
}
