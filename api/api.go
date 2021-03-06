// Copyright 2015 Features authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/albertoleal/features"
	"github.com/albertoleal/features/engine"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/tylerb/graceful"
	"golang.org/x/net/context"
)

const (
	DEFAULT_PORT    = ":8000"
	DEFAULT_TIMEOUT = 10 * time.Second
)

type Api struct {
	ng     engine.Engine
	router *Router
}

func NewApi(ng engine.Engine) *Api {
	api := &Api{router: NewRouter(), ng: ng}

	api.router.NotFoundHandler(http.HandlerFunc(api.notFoundHandler))
	api.router.AddHandler(RouterArguments{Path: "/", Methods: []string{"GET"}, Handler: homeHandler})
	ctx := context.Background()

	ffs := features.New(api.ng)

	createHandler := httptransport.NewServer(
		ctx,
		CreateEndpoint(ffs),
		decodeFeatureFlagRequest,
		encodeResponse,
		httptransport.ServerErrorEncoder(handleErrorEncoder),
	)
	api.router.AddHandler(RouterArguments{Path: "/features", Methods: []string{"POST"}, HandlerNormal: createHandler})

	updateHandler := httptransport.NewServer(
		ctx,
		UpdateEndpoint(ffs),
		decodeFeatureFlagRequest,
		encodeResponse,
		httptransport.ServerErrorEncoder(handleErrorEncoder),
	)
	api.router.AddHandler(RouterArguments{Path: "/features/{feature_key}", Methods: []string{"PUT"}, HandlerNormal: updateHandler})

	deleteHandler := httptransport.NewServer(
		ctx,
		DeleteEndpoint(ffs),
		decodeFeatureFlagQueryString,
		encodeResponse,
		httptransport.ServerErrorEncoder(handleErrorEncoder),
	)
	api.router.AddHandler(RouterArguments{Path: "/features/{feature_key}", Methods: []string{"DELETE"}, HandlerNormal: deleteHandler})

	findHandler := httptransport.NewServer(
		ctx,
		FindEndpoint(ffs),
		decodeFeatureFlagQueryString,
		encodeResponse,
		httptransport.ServerErrorEncoder(handleErrorEncoder),
	)
	api.router.AddHandler(RouterArguments{Path: "/features/{feature_key}", Methods: []string{"GET"}, HandlerNormal: findHandler})

	listHandler := httptransport.NewServer(
		ctx,
		ListEndpoint(ffs),
		decodeFeatureFlagQueryString,
		encodeResponse,
		httptransport.ServerErrorEncoder(handleErrorEncoder),
	)
	api.router.AddHandler(RouterArguments{Path: "/features", Methods: []string{"GET"}, HandlerNormal: listHandler})

	validateHandler := httptransport.NewServer(
		ctx,
		ValidateEndpoint(ffs),
		decodeValidationRequest,
		encodeResponse,
		httptransport.ServerErrorEncoder(handleErrorEncoder),
	)
	api.router.AddHandler(RouterArguments{Path: "/features", Methods: []string{"PUT"}, HandlerNormal: validateHandler})

	return api
}

func (api *Api) Handler() http.Handler {
	return api.router.Handler()
}

func (api *Api) Run() {
	fmt.Printf("Features is now ready to accept connections on port %s.", DEFAULT_PORT)
	graceful.Run(DEFAULT_PORT, DEFAULT_TIMEOUT, api.Handler())
}

func homeHandler(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "Hello Features!")
}

func (api *Api) notFoundHandler(rw http.ResponseWriter, r *http.Request) {
	handleErrorEncoder(rw, NewNotFoundError(ErrNotFound))
}
