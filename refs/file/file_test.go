package file

import (
	"errors"
	"io/fs"
	"reflect"
	"testing"
	"testing/fstest"
)

type MockFailingFs struct{
}

func (fs MockFailingFs) Open(name string) (fs.File, error) {
	return nil, errors.New("always failing")
}

func TestPosts(t *testing.T){
	const (
		firstBody = `Title: Post 1
Description: Description 1
Tags: tdd, go
---
Hello
World`
		secondBody = `Title: Post 2
Description: Description 2
Tags: tdd, go`
	)
	fs := fstest.MapFS{
		"hello world.md": {Data: []byte(firstBody)},
		"hello world2.md": {Data: []byte(secondBody)},
	}

	posts, err := postsFromFs(fs)

	if err != nil {
		t.Fatal(err)
	}

	assertPost(t, posts[0], Post{
		Title:       "Post 1",
		Description: "Description 1",
		Tags:        []string{"tdd", "go"},
		Body: `Hello
World`,
	})
}

func assertPost(t *testing.T, got Post, want Post) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}
