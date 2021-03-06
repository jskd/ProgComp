package share

import (
	"container/list"
	"fmt"
	"os"
	"reflect"
	"testing"
)

func AssertEqual(t *testing.T, a interface{}, b interface{}, message string) {
	if reflect.DeepEqual(a, b) {
		return
	} else {
		message = fmt.Sprintf("%s: %v != %v", message, a, b)
	}
	t.Fatal(message)
}

//TODO: To make general type double array [][]
func AssertArrayEqual(t *testing.T, input [][]string, expected [][]string, message string) {
	if len(input) != len(expected) {
		message = fmt.Sprintf("%s: Input and expect array length are different: %d != %d", message, len(input), len(expected))
	} else {
		var total int
		total = len(expected)
		for i := 0; i < (len(input)); i++ {
			for j := 0; j < (len(input)); j++ {
				if input[i][j] != expected[i][j] {
					message = fmt.Sprintf("%s: Input[%d][%d] != Expected[%d][%d]\nInput : %v\nExpect: %v", message, i, j, i, j, input[i][j], expected[i][j])
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

// Returns the temporary directory used by the program ws.  The
// directory name returned by this function ends with a path separator.
func TempDir() string {
	return AddPathSeparator(AddPathSeparator(os.TempDir()) + "ws")
}

func ArrayFromList(l *list.List) []interface{} {
	res := make([]interface{}, l.Len())
	for e, i := l.Front(), 0; e != nil; e = e.Next() {
		res[i] = e.Value
		i++
	}
	return res
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
