package test

import (
	"testing"

	"github.com/albertoleal/features/engine"
	. "gopkg.in/check.v1"
)

//Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

type EngineSuite struct {
	Engine engine.Engine
}

func (s *EngineSuite) TestFeatureFlagCRUD(c *C) {
	ffs, err := s.Engine.GetFeatureFlags()
	c.Check(err, IsNil)
	c.Assert(ffs, DeepEquals, []engine.FeatureFlag{})

	feature := engine.FeatureFlag{Key: "Feature A"}
	c.Check(s.Engine.UpsertFeatureFlag(feature), IsNil)

	ffs, err = s.Engine.GetFeatureFlags()
	c.Check(err, IsNil)
	c.Assert(ffs, DeepEquals, []engine.FeatureFlag{feature})

	ffk := engine.FeatureFlagKey{Key: "Feature A"}

	ff, err := s.Engine.GetFeatureFlag(ffk)
	c.Check(err, IsNil)
	c.Assert(ff, DeepEquals, &feature)

	c.Assert(s.Engine.DeleteFeatureFlag(ffk), IsNil)

	ffs, err = s.Engine.GetFeatureFlags()
	c.Check(err, IsNil)
	c.Assert(ffs, DeepEquals, []engine.FeatureFlag{})
}

func (s *EngineSuite) TestGetFeatureFlagWhenNotFound(c *C) {
	ffk := engine.FeatureFlagKey{Key: "Non Existing Feature Flag"}
	ff, err := s.Engine.GetFeatureFlag(ffk)
	c.Assert(err, DeepEquals, &engine.NotFoundError{})
	c.Check(ff, IsNil)
}

func (s *EngineSuite) TestDeleteFeatureFlagWhenNotFound(c *C) {
	ffk := engine.FeatureFlagKey{Key: "Non Existing Feature Flag"}
	err := s.Engine.DeleteFeatureFlag(ffk)
	c.Assert(err, DeepEquals, &engine.NotFoundError{})
}
