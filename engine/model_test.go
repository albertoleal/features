package engine

import (
	"testing"

	. "gopkg.in/check.v1"
)

type S struct {
}

var _ = Suite(&S{})

//Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

func (s *S) TestNewFeatureFlag(c *C) {
	ff, err := NewFeatureFlag("Feature A", true, []*User{})
	c.Check(err, IsNil)
	c.Assert(ff.Key, Equals, "Feature A")
	c.Assert(ff.Enabled, Equals, true)
}

func (s *S) TestNewFeatureFlagWithInvalidFields(c *C) {
	ff, err := NewFeatureFlag("", true, []*User{})
	c.Check(ff, IsNil)
	c.Check(err, Not(IsNil))
}

func (s *S) TestContainsUser(c *C) {
	user := &User{Id: "alice@example.org"}
	ff, err := NewFeatureFlag("Feature A", true, []*User{user})
	c.Check(err, IsNil)
	c.Assert(ff.ContainsUser(user), Equals, true)
}

func (s *S) TestContainsUserNotFound(c *C) {
	ff, err := NewFeatureFlag("Feature A", true, []*User{})
	c.Check(err, IsNil)
	user := &User{Id: "alice@example.org"}
	c.Assert(ff.ContainsUser(user), Equals, false)
}
