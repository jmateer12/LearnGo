// https://www.youtube.com/watch?v=c8H0w4yBL10&t=32s video used to further learn structs
package main

import (
	"encoding/json"
	"fmt"
)

// Type defines we want a new named type and UPPERCASE name means public to other modules, whereas lowercase means private and can't be used in other modules
// We use structs to create groupings of related data for more complicated data types
type Employee struct {

	// Tags in Go allow us go give data a sort of metadata we would like to convert to such as json and give names to the keys we want to use
	Name     string `json:"name"`
	Age      int    `json:"age"`
	IsRemote bool   `json:"isRemote"`
	Address
}

type Address struct {
	Street string `json:"myStreet"`
	City   string `json:"myCity"`
}

// To make this function a method of Employee we use the () at teh beginning to tell it to use the Employee struct
// Go uses the convention of naming the struct variable lowercase of the first letter of the struct so in this case e
// Lastly to modify the struct property of the variable callings this function we must use * which is a pointer to the struct calling
// If you don't use a pointer Go just passes by value not by reference to the struct calling the method
func (e *Employee) updateName(newName string) {
	e.Name = newName
}

func (a Address) printAddress() {
	fmt.Println("Address:", a.Street, a.City)
}

func main() {
	address := Address{
		Street: "123 Main Street",
		City:   "New York",
	}

	// This creates employee1 of type Employee and initialized in the curly braces the values
	// Initializing properties in a struct requires using field: value, you must include the comma
	employee1 := Employee{
		Name:     "Alice",
		Age:      30,
		IsRemote: true,

		// Using other types nested inside other structs is an example of composition
		Address: address,
	}

	employee1.printAddress()

	employee1.updateName("Bob")

	// Printing to console the properties of employee1 struct
	// structs have the advantage that they allows us to have random access of properties within them
	fmt.Println("Employee name:", employee1.Name)
	fmt.Println("Employee age:", employee1.Age)
	fmt.Println("Is employee remote:", employee1.IsRemote)

	// Here we have an anonymous struct which is a struct without a name this can simplify struct usage if we don't want to reuse them
	// One thing about anonymous structs is that they must be defined after being initialized using the second set of curly braces {}
	job := struct {
		title  string
		salary int
	}{
		title:  "Software Engineer",
		salary: 100000,
	}

	// To access struct fields we use . notation using variable.field much like we might do with objects in OOP languages
	fmt.Println("Job:", job.title)
	fmt.Println("Salary: $", job.salary)

	// employeePtr is just referencing the address of employee1 as denoted by &
	// We can use employeePtr to mutate the age of employee1 because it's just referencing its memory address
	employeePtr := &employee1
	employeePtr.Age = 31

	// The json.Marshal() function serializes the data of employee1
	// jsonData contains the bytes of the serialized data so we use string() to cast to a string type so we can understand it
	jsonData, _ := json.Marshal(employee1)
	fmt.Println(string(jsonData))
}
