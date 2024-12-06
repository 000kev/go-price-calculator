package main

import (
	"example.kyg/tax_calculator/tax_calculator"
	"fmt"
	"example.kyg/tax_calculator/file_manager"
	"example.kyg/tax_calculator/cmd_manager"
	"example.kyg/tax_calculator/io_manager"
)

type iom io_manager.IO_manager

const inputFileName string = "inputs.txt"
const outputFileName string = "results.json"
var rates = []float64{0, 0.25, 0.5}

func main() {
	// prices := []float64{10,20,30,40,50}
	var choice int
	prices := []float64{}

	fmt.Println("_____TAX CALCULATOR_____")
	fmt.Println("\nWould you like to load inputs from the command-line (1) or a file (2)?")
	fmt.Scan(&choice)

	for !validate(choice, &prices) {
		fmt.Scan(&choice)
	}

	t, err := tax_calculator.New(prices, rates)
	if err != nil {
		fmt.Println(err)
		return
	}

	t.Process()
	writeSuccess := t.Write(outputFileName)
	if writeSuccess != nil {
		fmt.Println("Something went wrong when saving the file")
		return
	}
	t.Display()

}

func validate(choice int, prices *[]float64) bool {
	if choice == 1 {
		cm := cmd_manager.New()
		*prices = load(cm)
		return true
	} else if choice == 2 {
		fm := file_manager.New(inputFileName, outputFileName)
		*prices = load(fm)
		return true
	} else {
		fmt.Println("Please provide a valid choice")
		return false
	}
}

func load(task iom) []float64 {
	return task.Load()
}
