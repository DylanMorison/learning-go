package main

// import "fmt"

// // List of functions that any variable that satisfies the animal type must have
// type Animal interface {
// 	Says() string
// 	NumberOfLegs() int
// }

// type Dog struct {
// 	Name  string
// 	Breed string
// }

// type Gorilla struct {
// 	Name          string
// 	Color         string
// 	NumberOfTeeth int
// }

// func main() {
// 	dog := Dog{
// 		Name:  "Samson",
// 		Breed: "Germain Sheppard",
// 	}

// 	PrintInfo(&dog)

// 	gorilla := Gorilla{
// 		Name:          "Jock",
// 		Color:         "grey",
// 		NumberOfTeeth: 38,
// 	}
// 	PrintInfo(&gorilla)
// }

// func PrintInfo(a Animal) {
// 	fmt.Println("This animal says", a.Says(), "and has", a.NumberOfLegs())
// }

// func (d *Dog) Says() string {
// 	return "woof"
// }

// func (d *Dog) NumberOfLegs() int {
// 	return 4
// }

// func (d *Gorilla) Says() string {
// 	return "woof"
// }

// func (d *Gorilla) NumberOfLegs() int {
// 	return 4
// }