package tool

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsStructEmpty(t *testing.T) {
	type Foo struct {
		Bar          string
		IntField     int
		BooleanField bool
	}
	foo := new(Foo)
	assert.Equal(t, true, IsStructEmpty(foo), "foo should be empty")
}
