// reflect.Tag.Get() Function

package main

import (
	"fmt"
	"reflect"
)

// Main function
func main() {
	type tag struct {
		val string `Value_1:"GeeksforGeeks" Value_2:"Best Platform :"`
	}

	src := tag{}
	st := reflect.TypeOf(src)
	field := st.Field(0)

	// use of Get method
	fmt.Println(field.Tag.Get("Value_2"), field.Tag.Get("Value_1"))

}

// reflect.TypeOf().Field().Tag.Get()
