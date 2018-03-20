package main

import (
	"errors"
	"fmt"
	"log"
	"strconv"
)

func main() {
	x, err := StrInverse("8")
	if err != nil {
		log.Printf("%s", err)
		return
	}
	fmt.Println(x)

}

func foo(n int) (int, error) {
	return 0, errors.New("Broken")
}

func test(a string) {
	fmt.Print(a)
}

func StrInverse(n string) (float64, error) {
	denom, err := strconv.ParseFloat(n, 64)
	if err != nil {
		return 0, fmt.Errorf("StrInver: %s %v", "Number is not a string", err)
	}
	if denom == 0 {
		return 0, errors.New("Can't divide by zero")
	}

	return 1 / float64(denom), nil
}
