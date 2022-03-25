// reflect.Set() Function

package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

// Main function
func main() {
	var s = struct{ foo int }{654}
	var i int

	rs := reflect.ValueOf(&s).Elem()
	rf := rs.Field(0)
	ri := reflect.ValueOf(&i).Elem()

	rf = reflect.NewAt(rf.Type(), unsafe.Pointer(rf.UnsafeAddr())).Elem()
	ri.Set(rf)
	rf.Set(ri)
	fmt.Println(rf)
	fmt.Println(ri)
}
