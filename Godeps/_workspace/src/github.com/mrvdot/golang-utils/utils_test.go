package utils

import (
	"reflect"
	"testing"
	. "launchpad.net/gocheck"
	"appengine/aetest"
)

type MySuite struct{}

var (
	_   = Suite(&MySuite{})
	ctx aetest.Context
)

// Hook up gocheck testing library to our usual testing tool
func Test(t *testing.T) {
	TestingT(t)
}

func (s *MySuite) TestGenerateSlug(c *C) {
	testString := "My awesome string"
	want := "my-awesome-string"
	slug := GenerateSlug(testString)
	c.Assert(slug, Equals, want)
}

func (s *MySuite) TestInChain(c *C) {
	chain := []string{"one", "two", "three"}
	yep := InChain("two", chain)
	c.Assert(yep, Equals, true)
	nope := InChain("four", chain)
	c.Assert(nope, Equals, false)
}

func (s *MySuite) TestIsEmpty(c *C) {
	// test string
	emptyStr := ""
	nonEmptyStr := "val"
	c.Assert(IsEmpty(reflect.ValueOf(emptyStr)), Equals, true)
	c.Assert(IsEmpty(reflect.ValueOf(nonEmptyStr)), Equals, false)
	emptyNum := 0
	nonEmptyNum := 5
	c.Assert(IsEmpty(reflect.ValueOf(emptyNum)), Equals, true)
	c.Assert(IsEmpty(reflect.ValueOf(nonEmptyNum)), Equals, false)
	var emptyPtr *string
	nonEmptyPtr := new(string)
	c.Assert(IsEmpty(reflect.ValueOf(emptyPtr)), Equals, true)
	c.Assert(IsEmpty(reflect.ValueOf(nonEmptyPtr)), Equals, false)
	emptySlice := []string{}
	nonEmptySlice := []string{"bob"}
	c.Assert(IsEmpty(reflect.ValueOf(emptySlice)), Equals, true)
	c.Assert(IsEmpty(reflect.ValueOf(nonEmptySlice)), Equals, false)
	emptyStruct := ApiResponse{}
	nonEmptyStruct := ApiResponse{Code: 200}
	c.Assert(IsEmpty(reflect.ValueOf(emptyStruct)), Equals, true)
	c.Assert(IsEmpty(reflect.ValueOf(nonEmptyStruct)), Equals, false)
}

func (s *MySuite) TestUpdate(c *C) {
	// should update in place
	original := &ApiResponse{
		Code:    404,
		Message: "Original Message",
	}
	data := &ApiResponse{
		Code:   200,
		Result: "good value",
	}
	Update(original, data)
	c.Assert(original, DeepEquals, &ApiResponse{
		Code:    200,
		Message: "Original Message",
		Result:  "good value",
	})
}
