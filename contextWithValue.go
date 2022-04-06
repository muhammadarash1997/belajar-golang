// Context sebagai Passing Value

// Hal paling sederhana dari context ialah kita dapat mengirim nilai atau data (object, string, dsb) melalui context dan nilai tersebut akan dapat diakses di seluruh turunan context yang terhubung. Misalnya,

// ctx := context.WithValue(parentContext, key, value)

//     parentContext : merupakan parent (induk) dari context tersebut,
//     key: merupakan key yang akan digunakan untuk mengambil data pada context, dan
//     value: merupakan nilai yang akan dikirim context dan semua turunannya.

// Konsep context seperti pohon (tree) yang bercabang-cabang.

package main

import (
	"context"
	"fmt"
)

func main() {
	// context.TODO() atau context.Background() sama saja, sama sama mengembalika empty Context
	
	contextParent := context.TODO() // Untuk membuat objek bertipe Context
	context.WithCancel()

	ctx1 := context.WithValue(contextParent, "key1", "hello world")
	ctx2 := context.WithValue(ctx1, "key2", "hello girls")
	ctx3 := context.WithValue(ctx2, "key3", "hello boys")
	ctx4 := context.WithValue(contextParent, "key4", "hello children")
	ctx5 := context.WithValue(ctx1, "key5", "hello Later")
	
	fmt.Println(ctx5.Value("key5")) // hello Later
	fmt.Println(ctx5.Value("key4")) // <nil>
	fmt.Println(ctx5.Value("key3")) // <nil>
	fmt.Println(ctx5.Value("key2")) // <nil>
	fmt.Println(ctx5.Value("key1")) // hello world
	fmt.Println("========")
	fmt.Println(ctx4.Value("key5")) // <nil>
	fmt.Println(ctx4.Value("key4")) // hello children
	fmt.Println(ctx4.Value("key3")) // <nil>
	fmt.Println(ctx4.Value("key2")) // <nil>
	fmt.Println(ctx4.Value("key1")) // <nil>
	fmt.Println("========")
	fmt.Println(ctx3.Value("key5")) // <nil>
	fmt.Println(ctx3.Value("key4")) // <nil>
	fmt.Println(ctx3.Value("key3")) // hello boys
	fmt.Println(ctx3.Value("key2")) // hello girls
	fmt.Println(ctx3.Value("key1")) // hello world
	fmt.Println("========")
	fmt.Println(ctx2.Value("key5")) // <nil>
	fmt.Println(ctx2.Value("key4")) // <nil>
	fmt.Println(ctx2.Value("key3")) // <nil>
	fmt.Println(ctx2.Value("key2")) // hello girls
	fmt.Println(ctx2.Value("key1")) // hello world
	fmt.Println("========")
	fmt.Println(ctx1.Value("key5")) // <nil>
	fmt.Println(ctx1.Value("key4")) // <nil>
	fmt.Println(ctx1.Value("key3")) // <nil>
	fmt.Println(ctx1.Value("key2")) // <nil>
	fmt.Println(ctx1.Value("key1")) // hello world
}