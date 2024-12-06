package cmd_manager

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CMDLineManager struct {}

func New() CMDLineManager {
	return CMDLineManager{}
}

func (cmd CMDLineManager) Load() []float64 {
	fmt.Println("Enter prices separated by a space (press enter to stop): ")

	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	pricesString := strings.Split(text, " ")

	prices := make([]float64, 0, len(pricesString))

	for _, val := range pricesString {
		fl, _ := strconv.ParseFloat(val, 64)
		// error checking if words are inputted

		prices = append(prices, fl)
	}

	return prices
}