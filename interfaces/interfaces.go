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
}

type Rectangle struct {
	width, height float64
}

// Implements Area() method for Rectangle
func (r Rectangle) Area() float64 {
	return r.width * r.height
}

type Circle struct {
	radius float64
}

// Implements Area() method for Circle
func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func calculateArea(s Shape) float64 {
	return s.Area()
}

func main() {
	//  This is how to define variables that contain struct data
	rect := Rectangle{width: 5, height: 4}
	circle := Circle{radius: 2}

	fmt.Println(
		"Rectangle area:",
		calculateArea(rect),
	)

	fmt.Println(
		"Circle area:",
		calculateArea(circle),
	)

	mysteryBox := interface{}(10)
	describeValue(mysteryBox)

	retrievedInt, ok := mysteryBox.(int)

	if ok {
		fmt.Println("Retrieved int:", retrievedInt)
	} else {
		fmt.Println("Value is not an integer")
	}
}

func describeValue(t interface{}) {
	fmt.Printf("Type: %T, Value: %v\n", t, t)
}
