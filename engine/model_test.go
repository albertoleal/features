package engine

import (
	"hash/crc32"
	"testing"

	. "gopkg.in/check.v1"
)

type S struct {
}

var _ = Suite(&S{})

//Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

func (s *S) TestNewFeatureFlag(c *C) {
	ff, err := NewFeatureFlag("Feature A", true, []*User{}, 0)
	c.Check(err, IsNil)
	c.Assert(ff.Key, Equals, "Feature A")
	c.Assert(ff.Enabled, Equals, true)
}

func (s *S) TestNewFeatureFlagWithInvalidFields(c *C) {
	ff, err := NewFeatureFlag("", true, []*User{}, 0)
	c.Check(ff, IsNil)
	c.Check(err, Not(IsNil))
}

func (s *S) TestContainsUser(c *C) {
	user := &User{Id: "alice@example.org"}
	ff, err := NewFeatureFlag("Feature A", true, []*User{user}, 0)
	c.Check(err, IsNil)
	c.Assert(ff.ContainsUser(user), Equals, true)
}

func (s *S) TestContainsUserNotFound(c *C) {
	ff, err := NewFeatureFlag("Feature A", true, []*User{}, 0)
	c.Check(err, IsNil)
	user := &User{Id: "alice@example.org"}
	c.Assert(ff.ContainsUser(user), Equals, false)
}

func (s *S) TestUserInPercentage(c *C) {
	ff, err := NewFeatureFlag("Feature A", true, []*User{}, 0)
	c.Check(err, IsNil)

	email := "alice@example.org"
	percentage := crc32.ChecksumIEEE([]byte(email)) % 100

	user := &User{Id: email}
	c.Assert(ff.UserInPercentage(user), Equals, false)

	ff, err = NewFeatureFlag("Feature A", true, []*User{}, percentage)
	c.Check(err, IsNil)
	c.Assert(ff.UserInPercentage(user), Equals, true)
}
