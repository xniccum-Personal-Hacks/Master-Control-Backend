package main

import (
	"fmt"
)

func add(x int, y int) (result int) {
	result = x * y
	return
}

func main() {
	fmt.Println(add(2, 3))
}
