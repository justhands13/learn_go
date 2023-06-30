package main

import (
	"fmt"

	r "github.com/justhands13/reverse_int/v2"
	s "github.com/justhands13/reverse_string/v2"
	tinytime "github.com/wagslane/go-tinytime"
)

func main() {
	fmt.Println("hello world")
	fmt.Println(tinytime.Now())
	fmt.Println(s.ReverseString("string desrever"))
	fmt.Println(r.ReverseInt(12938))
}
