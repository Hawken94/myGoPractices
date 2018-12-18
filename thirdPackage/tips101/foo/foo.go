// foo.go
package foo

import (
	"fmt"
	"unsafe"
)

type Config struct {
	_    struct{}
	Name string
	Size int
}

func P() {
	// conf := Config{0, "haha", 123} // doesn't compile
	conf := Config{Name: "bar", Size: 123} // compile okay
	// spew.Dump(conf)
	fmt.Println(conf)
	// fmt.Println(runtime.Version())
	fmt.Println("size p:", unsafe.Sizeof(conf))
}
