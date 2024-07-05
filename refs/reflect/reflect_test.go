package reflect

import (
	"reflect"
	"testing"
)

type Person struct {
	Name string
	Profile Profile
}

type Profile struct{
	Age int
	City string
}

func TestWalk(t *testing.T) {
	cases := []struct{
		Name string
		Input interface{}
		ExpectedCalls []string
	}{
		{
			"struct with one string field",
			struct{
				Name string
			}{"Vas"},
			[]string{"Vas"},
		},
		{
			"struct with two string fields",
			struct{
				Name string
				City string
			}{"Vas", "London"},
			[]string{"Vas","London"},
		},
		{
			"struct with non string field",
			struct{
				Name string
				Age int
			}{"Vas", 23},
			[]string{"Vas"},
		},
		{
			"struct with nested field",
			Person{"Vas", Profile{23, "London"}},
			[]string{"Vas", "London"},
		},
		{
			"pointer to thing",
			&Person{"Vas", Profile{23, "London"}},
			[]string{"Vas", "London"},
		},
		{
			"slices",
			[]Profile{
				{33, "Surrey"},
				{23, "Kent"},
			},
			[]string{"Surrey", "Kent"},
		},
		{
			"arrays",
			[2]Profile{
				{33, "Surrey"},
				{23, "Kent"},
			},
			[]string{"Surrey", "Kent"},
		},
		{
			"maps",
			map[string]string{
				"Vas": "Surrey",
				"Geenie": "Kent",
			},
			[]string{"Surrey", "Kent"},
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

	t.Run("with channels", func(t *testing.T) {
		profChannel := make(chan Profile)

		go func() {
			profChannel <- Profile{33, "Surrey"}
			profChannel <- Profile{23, "Kent"}
			close(profChannel)
		}()

		var got []string
		want := []string{"Surrey", "Kent"}

		walk(profChannel, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("with func", func(t *testing.T) {
		profFunc := func() (Profile, Profile) {
			return Profile{33, "Surrey"}, Profile{23, "Kent"}
		}

		var got []string
		want := []string{"Surrey", "Kent"}

		walk(profFunc, func(input string) {
			got = append(got, input)
		})


		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}