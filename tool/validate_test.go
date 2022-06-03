package tool

import (
	"testing"
)

func TestIsStructEmpty(t *testing.T) {
	// table test
	type test struct {
		name string
		in   interface{}
		out  bool
	}
	type Foo struct {
	}
	type Bar struct {
		Name string
	}
	type Baz struct {
		Int32Slice []int32
	}
	type Qux struct {
		Foo Foo
	}
	type Quux struct {
		Qux Qux
	}
	str := ""
	tests := []test{
		{
			name: "empty struct",
			in:   Foo{},
			out:  true,
		},
		{
			name: "not empty struct",
			in: Bar{
				Name: "foo",
			},
			out: false,
		},
		{
			name: "empty struct pointer",
			in:   new(Foo),
			out:  true,
		},
		{
			name: "struct with empty struct",
			in: Qux{
				Foo: Foo{},
			},
			out: true,
		},
		{
			name: "empty single string field struct",
			in: Bar{
				Name: "",
			},
			out: true,
		},
		{
			name: "empty single int field struct",
			in: struct {
				Name int
			}{
				Name: 0,
			},
			out: true,
		},
		{
			name: "empty single float field struct",
			in: struct {
				Name float64
			}{
				Name: 0,
			},
			out: true,
		},
		{
			name: "empty single uint field struct",
			in: struct {
				Name uint
			}{
				Name: 0,
			},
			out: true,
		},
		{
			name: "empty single map field struct",
			in: struct {
				Name map[string]string
			}{
				Name: map[string]string{},
			},
			out: true,
		},
		{
			name: "empty single slice field struct",
			in: &Baz{
				Int32Slice: []int32{},
			},
			out: true,
		},
		{
			name: "empty struct with nil map",
			in: struct {
				Name map[string]string
			}{
				Name: nil,
			},
			out: true,
		},
		{
			name: "empty struct with nil slice",
			in: struct {
				Name []string
			}{
				Name: nil,
			},
			out: true,
		},
		{
			name: "struct with slice which it's elements is empty struct",
			in: struct {
				Name []Foo
			}{
				Name: []Foo{{}},
			},
			out: true,
		},
		{
			name: "struct embed struct which the embed struct is nil",
			in: struct {
				Name *Foo
			}{
				Name: nil,
			},
			out: true,
		},
		{
			name: "struct embed struct which the embed struct is empty",
			in: struct {
				Name *Quux
			}{
				Name: &Quux{
					Qux: Qux{
						Foo: Foo{},
					},
				},
			},
			out: true,
		},
		{
			name: "struct embed struct pointer which the embed struct is empty",
			in: struct {
				Name *string
			}{
				Name: &str,
			},
			out: true,
		},
		{
			name: "struct with map which it's elements is empty string",
			in: struct {
				Name map[string]string
			}{
				Name: map[string]string{"": ""},
			},
			out: true,
		},
		{
			name: "struct with map which it's elements is empty struct",
			in: struct {
				Name map[string]Foo
			}{
				Name: map[string]Foo{"foo": {}},
			},
			out: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsStructEmpty(tt.in); got != tt.out {
				t.Errorf("IsStructEmpty() = %v, want %v", got, tt.out)
			}
		})
	}
}
