package injection

import (
	"bytes"
	"fmt"
)

func Greet(b *bytes.Buffer, name string) {
	fmt.Fprintf(b, "Hello, %s", name)
}