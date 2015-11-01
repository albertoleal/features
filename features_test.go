package features_test

import (
	"testing"

	"github.com/albertoleal/features"
	"github.com/albertoleal/features/engine"
	"github.com/albertoleal/features/engine/memory"
	. "gopkg.in/check.v1"
)

type S struct {
	Features *features.Features
}

var _ = Suite(&S{})

//Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

func (s *S) SetUpTest(c *C) {
	s.Features = features.New(memory.New())
}

func (s *S) TestAddFeature(c *C) {
	key := "Feature Key"
	feature := engine.FeatureFlag{
		Key:     key,
		Enabled: true,
	}
	s.Features.AddFeature(feature)
	c.Assert(s.Features.IsActive(key), Equals, true)
}

func (s *S) TestIsActive(c *C) {
	// Invalid Key
	key := "Feature Key"
	c.Assert(s.Features.IsActive(key), Equals, false)

	//  Enabled
	feature := engine.FeatureFlag{
		Key:     key,
		Enabled: true,
	}
	s.Features.AddFeature(feature)
	c.Assert(s.Features.IsActive(key), Equals, true)

	// Disabled
	feature = engine.FeatureFlag{
		Key:     key,
		Enabled: false,
	}
	s.Features.AddFeature(feature)
	c.Assert(s.Features.IsActive(key), Equals, false)
}

func (s *S) TestIsInactive(c *C) {
	// Invalid Key
	key := "Feature Key"
	c.Assert(s.Features.IsInactive(key), Equals, true)

	//  Disabled
	feature := engine.FeatureFlag{
		Key:     key,
		Enabled: false,
	}
	s.Features.AddFeature(feature)
	c.Assert(s.Features.IsInactive(key), Equals, true)

	// Enabled
	feature = engine.FeatureFlag{
		Key:     key,
		Enabled: true,
	}
	s.Features.AddFeature(feature)
	c.Assert(s.Features.IsInactive(key), Equals, false)
}

func (s *S) TestWith(c *C) {
	var status bool = true
	key := "Feature Key"
	s.Features.With(key, func() {
		status = false
	})
	c.Assert(status, Equals, true)

	// Set the Feature Flag as enabled
	feature := engine.FeatureFlag{
		Key:     key,
		Enabled: true,
	}
	s.Features.AddFeature(feature)

	s.Features.With(key, func() {
		status = false
	})
	c.Assert(status, Equals, false)
}

func (s *S) TestWithout(c *C) {
	var status bool
	s.Features.Without("Invalid Key", func() {
		status = true
	})
	c.Assert(status, Equals, true)
}
