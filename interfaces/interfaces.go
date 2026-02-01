//https://www.youtube.com/watch?v=SX1gT5A9H-U video for learning interfaces in Golang
//at 6:08

package main

import (
	"fmt"
	"math"
)

// An interface is like a contract
// Interfaces define high level organization of methods without business logic
type Shape interface {
	Area() float64
	Abc()
}

type Rectangle struct {
	width, height float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func main() {
	fmt.Print("Hello")
}
