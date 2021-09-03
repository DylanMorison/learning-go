package main

import "fmt"

func main() {
	x := 7
	inc(&x)
	fmt.Println(x)
} 

func inc(x *int) {
	*x++
}

// func sum(x int, y int) int {
// 	return x + y
// }

// //! map
// func main() {
// 	vertices := make(map[string]int)

// 	vertices["triangle"] = 2
// 	vertices["square"] = 4
// 	vertices["dodeas"] = 12

// 	fmt.Println(vertices["triangle"])
// }

// //! map
// func main() {
// 	vertices := make(map[string]int)

// 	vertices["triangle"] = 2
// 	vertices["square"] = 4
// 	vertices["dodeas"] = 12

// 	fmt.Println(vertices["triangle"])
// }
//! slice
// func main() {
// 	a := []int{}
// 	a = append(a, 13)
// 	fmt.Println(a)
// }
//! Arr
// func main() {
// 	a := [6]int{}
// 	a = append(a, 13)
// 	fmt.Println(a)
// }
 