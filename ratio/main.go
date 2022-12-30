package main

import (
	"fmt"

	dectofrac "github.com/av-elier/go-decimal-to-rational"
)

func main() {
	fmt.Println("Ratio compression placeholder")
	fmt.Println(dectofrac.NewRatP(0.6666, 0.01).String())
}
