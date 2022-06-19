// Golang program to illustrate
// reflect.Elem() Function

package main

import (
	"fmt"
	"io"
	"os"
	"reflect"
)

// Main function
func main() {

	//use of Elem() method
	writerType := reflect.TypeOf((*io.Writer)(nil)).Elem()

	fileType := reflect.TypeOf((*os.File)(nil))
	fmt.Println(fileType.Implements(writerType))
}
