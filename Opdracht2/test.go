package main

import (
	"fmt"
	"math"
	"os"
)

var distanceM int
var rotationHours float64

func main() {
	asknumbers()
	speedcalculation()

}

func speedcalculation() {
	var circumference float64 = 2 * math.Pi * float64(distanceM)
	speed := circumference / rotationHours

	result := fmt.Sprintf("The speed of the planet is: %.2f km/h\n", speed)

	saveToFile(result)
	fmt.Print(result)
}

func asknumbers() {

	fmt.Println("What is the km distance of the planet from the sun? (in Millions)")
	_, err := fmt.Scanln(&distanceM)
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

func saveToFile(data string) {
	file, err := os.Create("result.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
	}
	defer file.Close()

	_, err = file.WriteString(data)
	if err != nil {
		fmt.Println("Error writing to file:", err)
	}
}
