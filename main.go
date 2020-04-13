package main

import (
	"fmt"
	. "sort/sortlogic"
)

func main() {
	a := []int{5, 3, 4, 2, 1, 6, 7}
	fmt.Println(Bubbling(a))
	fmt.Println(Insert(a))
	fmt.Println(Shell(a))
	fmt.Println(Stack(a))
	fmt.Println(Quick(a))
	fmt.Println(Merge(a))
	fmt.Println(Count(a))
	fmt.Println(Bucket(a))
	fmt.Println(Radix(a))
	var stop int
	fmt.Scanln(&stop)

}
