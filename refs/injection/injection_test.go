package injection

import (
	"bytes"
	"testing"
)

func TestInj(t *testing.T){
	t.Run("write to buffer", func(t *testing.T) {
		buffer := bytes.Buffer{}
		Greet(&buffer, "James")

		got := buffer.String()
		want := "Hello, James"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})	
}