package main

import (
	"geektrust/calculation"
	"geektrust/input"
)

func main() {
	userInput := input.FormattedInput()
	calculation.Calculation(userInput)

}
