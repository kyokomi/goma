package debuglog

import "fmt"

// Println Debug Println
func Println(v ...interface{}) {
	// SliceInsert (https://code.google.com/p/go-wiki/wiki/SliceTricks)
	v = append(v[:0], append([]interface{}{"[goma]", "[debug]"}, v[0:]...)...)

	fmt.Println(v...)
}
