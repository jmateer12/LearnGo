package main

import "fmt"

func main() {

	calculateGas()

}

func calculateGas() {
	var efficiency int = 0
	var miles int = 0
	var trips int = 0
	var price float32 = 0.0

	fmt.Print("How many miles per gallon do you get out of your car (on the low end): ")
	fmt.Scan(&efficiency)
	fmt.Print("How many miles do you drive to your destination each day: ")
	fmt.Scan(&miles)
	fmt.Print("How many trips do you make per week: ")
	fmt.Scan(&trips)
	fmt.Print("How much does a gallon of gas cost in your area? $")
	fmt.Scan(&price)

	fmt.Print("\nOver the course of your semester you will spend $", (float32((trips*16)*miles/efficiency) * price), "\n\n")
}
