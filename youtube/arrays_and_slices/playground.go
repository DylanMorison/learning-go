package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5}
	b := a[1:]
	fmt.Println(b)
}

// func main() {
// 	a := []int{}
// 	fmt.Println(a)      //[]
// 	fmt.Println(len(a)) //0
// 	fmt.Println(cap(a)) //0
// 	a = append(a, 1)
// 	fmt.Println(a)      //[1]
// 	fmt.Println(len(a)) //1
// 	fmt.Println(cap(a)) //1
// 	a = append(a, 2, 3, 4, 5)
// 	fmt.Println(a)      //[1 2 3 4 5]
// 	fmt.Println(len(a)) //5
// 	fmt.Println(cap(a)) //6
// }

// func main() {
// 	a := make([]int, 3, 100)
// 	fmt.Println(a)      //[0,0,0]
// 	fmt.Println(len(a)) //3
// 	fmt.Println(cap(a)) //100
// }

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	b := a[:]
	c := a[3:]
	d := a[:6]
	e := a[3:6]
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
}

func main() {
	a := []int{1, 2, 3}
	fmt.Println(len(a)) //3
	fmt.Println(cap(a)) //3
}

func main() {
	a := []int{1, 2, 3}
	// b := &a
	// b[1] = 5
	fmt.Println(a)
	// fmt.Println(b)
}

// func main() {
// 	a := [...]int{1, 2, 3}
// 	b := &a
// 	b[1] = 5
// 	fmt.Println(a)
// 	fmt.Println(b)
// }

// func main() {
// 	a := [...]int{1, 2, 3}
// 	b := a
// 	b[1] = 5
// 	fmt.Println(a)
// 	fmt.Println(b)
// }

// func main() {
// 	var identityMatrix [3][3]int
// 	identityMatrix[0] = [3]int{1, 0, 0}
// 	identityMatrix[1] = [3]int{0, 1, 0}
// 	identityMatrix[2] = [3]int{0, 0, 1}
// 	fmt.Println(identityMatrix)
// }

// func main() {
// 	var identityMatrix [3][3]int = [3][3]int{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}}

// 	fmt.Println(identityMatrix)

// }

// func main() {
// 	var students [3]string
// 	fmt.Printf("Students: %v\n", students)
// 	students[0] = "Lisa"
// 	students[1] = "Dylan"
// 	students[2] = "Tom"
// 	fmt.Printf("Students: %v\n", students)
// 	fmt.Printf("Students[1]: %v\n", students[1])
// 	fmt.Printf("len(students): %v\n", len(students))
// }

// func main() {
// 	var students [3]string
// 	fmt.Printf("Students: %v\n", students)
// 	students[0] = "Lisa"
// 	fmt.Printf("Students: %v", students)
// }

// func main() {
// 	// create an array called "firstArray" that has length 3
// 	// and holds integers
// 	firstArray := [3]int{1, 2, 3}
// 	fmt.Printf("firstArray: %v, type: %T", firstArray, firstArray)
// }
