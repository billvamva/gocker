package maps

import (
	"testing"
)

func TestMap(t *testing.T){

	assertSearch := func(t testing.TB, key string, got string, want string) {
		if got != want {
			t.Errorf("got %q want %q, given %q", got, want, key)
		}
	}

	assertErr := func(t testing.TB, got error, want error) {
		if got != want {
			t.Errorf("got %q want %q.", got.Error(), want.Error())
		}
	}
	t.Run("known word", func(t *testing.T) {
		dictionary := Dictionary{"test":"this is a dictionary"}

		got, _ := dictionary.Search("test")
		want := "this is a dictionary"

		assertSearch(t, "test", got, want)

	})

	t.Run("unknown word", func(t *testing.T) {
		dictionary := Dictionary{"test":"this is a dictionary"}

		_, err := dictionary.Search("unknown")
		want := ErrUnknownKey

		assertErr(t,err,want)

	})

	t.Run("add key", func(t *testing.T) {
		dictionary := Dictionary{"test":"this is a dictionary"}
		key := "new"
		dictionary.Add(key, "added new value")

		got, _ := dictionary.Search("new")
		want := "added new value"

		assertSearch(t, key, got, want)
	})

	t.Run("update value", func(t *testing.T) {
		dictionary := Dictionary{"test":"this is a dictionary"}
		newValue := "this value was updated"
		dictionary.Update("test", newValue)

		got, _ := dictionary.Search("test")

		assertSearch(t, "test", got, newValue)
	})
	
	t.Run("update value", func(t *testing.T) {
		dictionary := Dictionary{"test":"this is a dictionary"}
		_ = dictionary.Delete("test")

		_, err := dictionary.Search("test")
		want := ErrUnknownKey

		assertErr(t, err, want)
	})

	
}