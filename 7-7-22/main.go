package main

import (
	"errors"
	"log"

	"github.com/DylanMorison/myNiceProgram/helpers"
	// "log"
)

const numPool = 10000

func CalculateValue(intChan chan int) {
	randomNumber := helpers.RandomNumber(numPool)
	intChan <- randomNumber
}

func Divide(x, y float32) (float32, error) {
	var result float32

	if y == 0 {
		return result, errors.New("cannot divide by zero")
	}
	result = x / y
	return result, nil

}

func main() {
	// intChan := make(chan int)
	// defer close(intChan)

	// go CalculateValue(intChan)

	// num := <-intChan
	// log.Println(num)
	log.Println(Divide(1, 0))

}
