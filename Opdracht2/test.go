package main

import "fmt"

var distanceKm int
var rotationHours int

func main() {
	asknumbers()
	speedcalculation()

}

func speedcalculation() {
	fmt.Println(distanceKm * rotationHours)
}

func asknumbers() {

	fmt.Println("What is the km distance of the planet from the sun? (in Millions)")
	_, err := fmt.Scanln(&distanceKm)
	if err != nil {
		fmt.Println("Error reading the first number. Please ensure you enter a valid number.")
		return
	}

	fmt.Println("In how many hours does the planet rotate around the sun?")
	_, err = fmt.Scanln(&rotationHours)
	if err != nil {
		fmt.Println("Error reading the first number. Please ensure you enter a valid number.")
		return
	}

}
