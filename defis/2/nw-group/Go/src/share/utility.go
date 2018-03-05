package share

import (
	"fmt"
	"os"
	"reflect"
	"testing"
)

func AssertEqual(t *testing.T, a interface{}, b interface{}, message string) {
	if reflect.DeepEqual(a, b) {
		return
	} else {
		message = fmt.Sprintf("%v != %v", a, b)
	}
	t.Fatal(message)
}

//TODO: To make general type double array [][]
func AssertArrayEqual(t *testing.T, input [][]string, expected [][]string, message string) {
	if len(input) != len(expected) {
		message = fmt.Sprintf("Input and expect array length are different: %d != %d", len(input), len(expected))
	} else {
		var total int
		total = len(expected)
		for i := 0; i < (len(input)); i++ {
			for j := 0; j < (len(input)); j++ {
				if input[i][j] != expected[i][j] {
					message = fmt.Sprintf("Input[%d][%d] != Expected[%d][%d]\nInput : %v\nExpect: %v", i, j, i, j, input[i][j], expected[i][j])
					t.Fatal(message)
				}
			}
			total -= 1
		}
		if total == 0 {
			return
		}
	}
	t.Fatal(message)
}

func AddPathSeparator(s string) string {
	return s + fmt.Sprintf("%c", os.PathSeparator)
}
