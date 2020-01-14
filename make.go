package main

import (
	"fmt"
)

func main() {

	slice := make([] int,0,4)
	slice = append(slice, 2)
	fmt.Println(cap(slice))
	fmt.Println(slice)
}
