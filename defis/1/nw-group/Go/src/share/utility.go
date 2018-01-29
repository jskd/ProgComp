package share

import (
	"fmt"
	"testing"
)

func AssertEqual(t *testing.T, a interface{}, b interface{}, message string) {
	if a == b {
		return
	} else {
		message = fmt.Sprintf("%v != %v", a, b)
	}
	t.Fatal(message)
}
