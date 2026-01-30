//https://www.youtube.com/watch?v=SX1gT5A9H-U video for learning interfaces in Golang

package main

import "fmt"

// An interface is like a contract
type Shape interface {
	Area() float64
}

func main() {
	fmt.Print("Hello")
}
