package tax_calculator

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type TaxCalculator struct {
	Input  []float64           `json:"input_prices"`
	Taxes  []float64           `json:"tax_rates"`
	Result map[string][]string `json:"tax_inclusive_prices"`
}

func (t TaxCalculator) Write(filename string) error {
	json, _ := json.Marshal(t)
	return os.WriteFile(filename, json, 0644)
}

func (t TaxCalculator) Display() {
	for key, value := range t.Result {
		fmt.Printf("Key: %.v Value: ", key)
		fmt.Println(value)
	}
}

func (t TaxCalculator) transform(tax float64) *[]string {
	a := make([]string, 0, len(t.Input))
	for _, val := range t.Input {
		a = append(a, fmt.Sprintf("%.2f", val+val*tax))
	}
	return &a
}

func (t TaxCalculator) Process() {
	// take input prices and generate the result
	// for each tax rate
	for _, tax := range t.Taxes {
		taxString := fmt.Sprintf("%.3f", tax)
		t.Result[taxString] = *t.transform(tax)
	}
}

func New(inputs []float64, taxes []float64) (*TaxCalculator, error) {
	if len(inputs) == 0 {
		return nil, errors.New("inputs must be non-empty")
	}
	if len(taxes) == 0 {
		return nil, errors.New("tax rates must be non-empty")
	}

	return &TaxCalculator{
		Input:  inputs,
		Taxes:  taxes,
		Result: make(map[string][]string, len(inputs)),
	}, nil
}
