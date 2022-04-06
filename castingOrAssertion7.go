package main

import (
    "fmt"
)

type User struct {
    ID int64
}

func main() {
    // 3 examples below is same
    
    x := (*User)(nil)
    var y *User = nil
    var z *User

    fmt.Println(x)
    fmt.Println(y)
    fmt.Println(z)
}
