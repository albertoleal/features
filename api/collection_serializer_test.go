package api_test

import (
	"github.com/albertoleal/features/api"
	"github.com/albertoleal/features/engine"
	. "gopkg.in/check.v1"
)

func (s *S) TestCollectionSerializer(c *C) {
	features := []*engine.FeatureFlag{
		&engine.FeatureFlag{Key: "feature-x", Enabled: true, Users: []*engine.User{}},
	}

	cs := &api.CollectionSerializer{
		Items: features,
		Count: len(features),
	}
	c.Assert(cs.Serializer(), Equals, `{"items":[{"enabled":true,"key":"feature-x","percentage":0}],"item_count":1}`)
}
