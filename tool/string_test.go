package tool

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTrimFields(t *testing.T) {
	type Foo struct {
		Name  string
		Age   int
		Email string
	}
	foo := &Foo{
		Name:  "   foo    ",
		Age:   10,
		Email: "bar@qux.com   ",
	}
	res := TrimFields(foo).(*Foo)
	assert.Equal(t, "foo", res.Name)
	assert.Equal(t, 10, res.Age)
	assert.Equal(t, "bar@qux.com", res.Email)
}
