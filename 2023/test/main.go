package main

import (
	_ "embed"
	"fmt"
)

func main() {
	b := [...]int{1, 2, 3}
	fmt.Println(b)
	Poop(b)
	fmt.Println(b)

}

func Poop(a [3]int) {
	fmt.Println(a)
	a[1] = 4
	fmt.Println(a)

}
