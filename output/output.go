// https://www.youtube.com/watch?v=NvfSa3XwvE0 using this video for learning about outputting in golang
// 10:13 in video

package main

import "fmt"

const (
	// 033 -> octal number (escape)
	// [
	Reset = "\033[0m"
	Red   = "\033[31m"
	Green = "\033[32m"
)

func main() {
	name := "Pete"
	age := 33

	type person struct {
		name string
		age  int
	}

	// How to print to console using colors and constants
	fmt.Println(Red + "This is Red" + Reset)
	fmt.Println(Green + "This is Green" + Reset)

	// How to print colors inline in golang
	fmt.Println("\033[33;1mThis is bright yellow" + Reset)

	// Standard way to print to console in golang and automatically adds a space between parameters
	// Adds a line after printing to console, adds a space
	fmt.Println("Name:", name, "Age:", age)

	// Prints something in a format, %s and %d are like placeholders
	// This is cleaner and easier to read and doesn't print a newline after printing to console

	/*
		%s -> String
		%d -> Integer
		%f -> Float (%.2f)
		%t -> Boolean, %T -> Type
		%v -> default format
		%+v -> struct
		%#v -> go syntax representation

	*/
	fmt.Printf("Name: %s, Age: %d\n", name, age)

	p := person{"Pete", 33}
	s := []int{1, 2, 3}

	fmt.Printf("Struct: %+v\n", p)
	fmt.Printf("Slice: %#v\n", s)
}
