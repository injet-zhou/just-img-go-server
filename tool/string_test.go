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
	type Bar struct {
	}
	bar := &Bar{}
	newBar := TrimFields(bar).(*Bar)
	assert.Equal(t, &Bar{}, newBar)
	// TODO fix it
	//var nilBar *Bar
	//assert.Nil(t, nilBar)
	//newNilBar := TrimFields(&nilBar)
	//assert.Nil(t, newNilBar)
	type Baz struct {
		Foo *Foo
	}
	baz := &Baz{
		Foo: foo,
	}
	newBaz := TrimFields(baz).(*Baz)
	assert.Equal(t, "foo", newBaz.Foo.Name)
	assert.Equal(t, 10, newBaz.Foo.Age)
	assert.Equal(t, "bar@qux.com", newBaz.Foo.Email)
}
