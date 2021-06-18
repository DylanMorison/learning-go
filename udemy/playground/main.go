package main

import "fmt"

func main() {
	name := "bill"

	namePointer := &name

	fmt.Println(&namePointer)
	printPointer(namePointer)
}

func printPointer(namePointer *string) {
	fmt.Println(&namePointer)
} // package main

// import "fmt"

// func main() {
// 	bill := "hi"
// 	fmt.Println(&bill)
// }

// package main

// import "fmt"

// func main() {
// 	mySlice := []string{"Hi", "There", "how", "are", "you!"}
// 	updateSlice(mySlice)
// 	fmt.Println(mySlice)
// }

// func updateSlice(s []string) {
// 	s[0] = "bye"
// }
