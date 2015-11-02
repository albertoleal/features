// Copyright 2015 Features authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package api

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"golang.org/x/net/context"
)

type Router struct {
	ctx      context.Context
	r        *mux.Router
	notFound http.Handler
	sub      map[string]*mux.Router
}

type RouterArguments struct {
	Handler       http.HandlerFunc
	HandlerNormal http.Handler
	Path          string
	PathPrefix    string
	Methods       []string
}

func NewRouter() *Router {
	return &Router{
		ctx: context.Background(),
		r:   mux.NewRouter(),
		sub: make(map[string]*mux.Router),
	}
}

func (router *Router) Handler() http.Handler {
	return router.r
}

func (router *Router) NotFoundHandler(h http.Handler) {
	router.notFound = h
	router.r.NotFoundHandler = h
}

func (router *Router) AddSubrouter(pathPrefix string) *mux.Router {
	s := mux.NewRouter()
	s.NotFoundHandler = router.notFound
	router.sub[pathPrefix] = s
	return s
}

func (router *Router) Subrouter(pathPrefix string) *mux.Router {
	return router.sub[pathPrefix]
}

func (router *Router) AddMiddleware(pathPrefix string, h http.Handler) {
	router.r.PathPrefix(pathPrefix).Handler(h)
}

func (router *Router) AddHandler(args RouterArguments) {
	var r *mux.Router

	if sub, ok := router.sub[args.PathPrefix]; ok {
		r = sub
	} else {
		r = router.r
	}

	var prefix, path string
	if args.PathPrefix != "" {
		prefix = fmt.Sprintf("/%s", strings.Trim(args.PathPrefix, "/"))
	}
	path = fmt.Sprintf("/%s", strings.Trim(args.Path, "/"))

	if args.Handler != nil {
		r.Methods(args.Methods...).Path(fmt.Sprintf("%s%s", prefix, path)).HandlerFunc(args.Handler)
	} else if args.HandlerNormal != nil {
		r.Methods(args.Methods...).Path(fmt.Sprintf("%s%s", prefix, path)).Handler(args.HandlerNormal)
	}
}
