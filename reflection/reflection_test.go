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
		{
			"Array",
			[2] Profile {
				{33, "Shanghai"},
				{35, "Beijing"},
			},
			[]string{"Shanghai", "Beijing"},
		},
		{
			"Maps",
			map[string]string {
				"Foo": "Bar",
				"Baz": "Boz",
			},
			[]string{"Bar", "Boz"},
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

	t.Run("with maps", func(t *testing.T){
		aMap := map[string] string{
			"Foo": "Bar",
			"Baz": "Boz",
		}

		var got []string
		walk(aMap, func(input string){
			got = append(got, input)
		})

		assertContains(t, got, "Bar")
		assertContains(t, got, "Boz")
	})

	t.Run("with channels", func(t *testing.T){
		aChannel := make(chan Profile)

		go func() {
			aChannel <- Profile{34, "Shanghai"}
			aChannel <- Profile{36, "Beijing"}
			close(aChannel)
		}()

		var got []string
		want := []string {"Shanghai", "Beijing"}
		walk(aChannel, func(input string){
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("with function", func(t *testing.T){
		aFunction := func() (Profile, Profile) {
			return Profile{34, "Shanghai"}, Profile{36, "Beijing"}
		}

		var got []string
		want := []string{"Shanghai", "Beijing"}

		walk(aFunction, func(input string){
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func assertContains(t *testing.T, haystack []string, needle string) {
	contains := false
	for _, x := range haystack {
		if x == needle {
			contains = true
			break
		}
	}

	if !contains {
		t.Errorf("expected %+v to contain %q but it didn't", haystack, needle)
	}
}