package ocaml

import (
	"share"
	"testing"
)

func TestOCamlCall(t *testing.T) {
	out := Execute("../../../OCaml/test.ml")
	share.AssertEqual(t, out, "Hello from OCaml!\n", "Incorrect OCaml return value.")
}
