package reflection

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
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
			"struct with one string field",
			struct {
				Name string
			}{"shashwat"},
			[]string{"shashwat"},
		},
		{
			"struct with two string fields",
			struct {
				Name string
				City string
			}{"Chris", "London"},
			[]string{"Chris", "London"},
		},
		{
			"struct with non string field",
			struct {
				Name string
				Age  int
			}{"Chris", 33},
			[]string{"Chris"},
		},
		{
			"nested fields",
			Person{
				"shashwat",
				Profile{33, "Japan"},
			},
			[]string{"shashwat", "Japan"},
		},
		{
			"pointers to things",
			&Person{
				"Chris",
				Profile{33, "London"},
			},
			[]string{"Chris", "London"},
		},
		{
			"slices",
			[]Profile{
				{33, "London"},
				{34, "Reykjavík"},
			},
			[]string{"London", "Reykjavík"},
		},
		{
			"arrays",
			[2]Profile{
				{33, "London"},
				{34, "Reykjavík"},
			},
			[]string{"London", "Reykjavík"},
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
	t.Run("with maps", func(t *testing.T) {
		aMap := map[string]string{
			"Cow":   "Moo",
			"Sheep": "Baa",
		}
		var got []string
		walk(aMap, func(input string) {
			got = append(got, input)
		})
		assert.Contains(t, got, "Moo", "expected %+v to contain %q but it didn't", got, "Moo")
		assert.Contains(t, got, "Baa", "expected %+v to contain %q but it didn't", got, "Baa")
	})

	t.Run("with channel", func(t *testing.T) {
		aChannel := make(chan Profile)
		go func() {
			aChannel <- Profile{33, "Berlin"}
			aChannel <- Profile{34, "Katowice"}
			close(aChannel)
		}()
		var got []string
		want := []string{"Berlin", "Katowice"}
		walk(aChannel, func(input string) {
			got = append(got, input)
		})
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
	t.Run("with function", func(t *testing.T) {
		aFunction := func() (Profile, Profile) {
			return Profile{33, "Berlin"}, Profile{34, "Katowice"}
		}
		var got []string
		want := []string{"Berlin", "Katowice"}
		walk(aFunction, func(input string) {
			got = append(got, input)
		})
		if !reflect.DeepEqual(got, want) {
			t.Errorf("want %v, got %v", got, want)
		}
	})
}
