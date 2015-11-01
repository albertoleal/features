package memory

import (
	"testing"

	"github.com/albertoleal/features/engine/test"
	. "gopkg.in/check.v1"
)

func TestMem(t *testing.T) {
	Suite(&test.EngineSuite{Engine: New()})
	TestingT(t)
}
