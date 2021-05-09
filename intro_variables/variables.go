package main

import "fmt"

func main() {
	n := 3.14
	n = 13.7e72
	n = 2.1e14
	fmt.Printf("%v, %T", n, n)
}

// func main() {
// 	a := 10             // 1010
// 	fmt.Println(a << 3) // 2^3 * 2^3 = 2^6
// 	fmt.Println(a >> 3) // 2^3 / 2^3 = 2^0
// }

// func main() {
// 	a := 10 // 1010
// 	b := 3  // 0011
// 	fmt.Println(a & b)  // 0010 AND
// 	fmt.Println(a | b)  // 1011 OR
// 	fmt.Println(a ^ b)  // 1001 EXCULSIVE OR
// 	fmt.Println(a &^ b) // 0100 AND NOT
// }

// func main() {
// 	n := 1 == 1
// 	m := 2 == 2
// 	var zero bool
// 	fmt.Printf("%v, %T\n", n, n)
// 	fmt.Printf("%v, %T\n", m, m)
// 	fmt.Printf("%v, %T\n", zero, zero)
// }
