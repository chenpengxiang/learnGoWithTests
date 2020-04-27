package main

import (
	"reflect"
	"testing"
)

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}

func TestWalk(t *testing.T) {

	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"Struct with one string field",
			struct {
				Name string
			}{"Peter"},
			[]string{"Peter"},
		},

		{
			"Struct with two string fields",
			struct {
				Name string
				City string
			}{"Peter", "Beijing"},
			[]string{"Peter", "Beijing"},
		},

		{
			"Struct with non string field",
			struct {
				Name string
				Age  int
			}{"Peter", 36},
			[]string{"Peter"},
		},

		{
			"Struct with Nested fields",
			Person{"Peter", Profile{36, "Beijing"}},
			[]string{"Peter", "Beijing"},
		},

		{
			"Pointers to things",
			&Person {
				"Peter",
				Profile{36, "Beijing"},
			},
			[]string{"Peter", "Beijing"}, 
		},
		{
			"Slices",
			[] Profile {
				{33, "Shanghai"},
				{35, "Beijing"},
			},
			[]string{"Shanghai", "Beijing"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(input string) {
				got = append(got, input)
			})
			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v, want %v", got, test.ExpectedCalls)
			}
		})
	}
}