package features_test

import (
	"hash/crc32"
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

func (s *S) TestSave(c *C) {
	key := "Feature Key"
	feature := engine.FeatureFlag{
		Key:     key,
		Enabled: true,
	}
	s.Features.Save(feature)
	active, err := s.Features.IsActive(key)
	c.Assert(active, Equals, true)
	c.Check(err, IsNil)
}

func (s *S) TestIsActive(c *C) {
	// Invalid Key
	key := "Feature Key"
	active, err := s.Features.IsActive(key)
	c.Assert(active, Equals, false)
	c.Check(err, Not(IsNil))

	//  Enabled
	feature := engine.FeatureFlag{
		Key:     key,
		Enabled: true,
	}
	s.Features.Save(feature)

	active, err = s.Features.IsActive(key)
	c.Assert(active, Equals, true)
	c.Check(err, IsNil)

	// Disabled
	feature = engine.FeatureFlag{
		Key:     key,
		Enabled: false,
	}
	s.Features.Save(feature)
	active, err = s.Features.IsActive(key)
	c.Assert(active, Equals, false)
	c.Check(err, IsNil)
}

func (s *S) TestIsActiveWithPercentage(c *C) {
	key := "Feature Key"
	feature := engine.FeatureFlag{
		Key:        key,
		Enabled:    true,
		Percentage: 50,
	}
	s.Features.Save(feature)

	active, err := s.Features.IsActive(key)
	c.Assert(active, Equals, false)
	c.Check(err, Not(IsNil))

}

func (s *S) TestIsInactive(c *C) {
	// Invalid Key
	key := "Feature Key"
	inactive, err := s.Features.IsInactive(key)
	c.Assert(inactive, Equals, true)
	c.Check(err, Not(IsNil))

	//  Disabled
	feature := engine.FeatureFlag{
		Key:     key,
		Enabled: false,
	}
	s.Features.Save(feature)

	inactive, err = s.Features.IsInactive(key)
	c.Assert(inactive, Equals, true)
	c.Check(err, IsNil)

	// Enabled
	feature = engine.FeatureFlag{
		Key:     key,
		Enabled: true,
	}
	s.Features.Save(feature)
	inactive, err = s.Features.IsInactive(key)
	c.Assert(inactive, Equals, false)
	c.Check(err, IsNil)
}

func (s *S) TestIsInactiveWithPercentage(c *C) {
	key := "Feature Key"
	feature := engine.FeatureFlag{
		Key:        key,
		Enabled:    false,
		Percentage: 50,
	}
	s.Features.Save(feature)

	inactive, err := s.Features.IsInactive(key)
	c.Assert(inactive, Equals, true)
	c.Check(err, Not(IsNil))

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
	s.Features.Save(feature)

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

func (s *S) TestUserHasAccessWhenTheFeatureIsInactive(c *C) {
	key := "Feature Key"
	email := "alice@example.org"

	feature, err := engine.NewFeatureFlag(key, false, []*engine.User{&engine.User{Id: email}}, 0)
	err = s.Features.Save(*feature)
	c.Check(err, IsNil)

	c.Assert(s.Features.UserHasAccess(key, email), Equals, true)
}

func (s *S) TestUserHasAccessWhenTheFeatureIsActive(c *C) {
	key := "Feature Key"
	email := "alice@example.org"

	feature, err := engine.NewFeatureFlag(key, true, []*engine.User{&engine.User{Id: email}}, 0)
	err = s.Features.Save(*feature)
	c.Check(err, IsNil)

	c.Assert(s.Features.UserHasAccess(key, email), Equals, true)
}

func (s *S) TestUserHasAccessWithPercentage(c *C) {
	key := "Feature Key"
	email := "alice@example.org"
	percentage := crc32.ChecksumIEEE([]byte(email)) % 100

	feature, err := engine.NewFeatureFlag(key, true, []*engine.User{}, percentage-1)
	err = s.Features.Save(*feature)
	c.Check(err, IsNil)
	c.Assert(s.Features.UserHasAccess(key, email), Equals, false)

	feature, err = engine.NewFeatureFlag(key, true, []*engine.User{}, percentage)
	err = s.Features.Save(*feature)
	c.Check(err, IsNil)
	c.Assert(s.Features.UserHasAccess(key, email), Equals, true)
}
