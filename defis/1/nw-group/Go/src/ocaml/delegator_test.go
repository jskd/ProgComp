package ocaml

import (
	"fmt"
	"testing"
)

func TestOCamlCall(t *testing.T) {
	fmt.Println("Testing OCaml Call...")
	out := Execute("../OCaml/test.ml")
	fmt.Printf("%q\n", out)
}
