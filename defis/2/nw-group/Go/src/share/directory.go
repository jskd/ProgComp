package share

import "os"

func Purge(path string) {
	if err := os.RemoveAll(path); err != nil {
		panic(err)
	}
}
