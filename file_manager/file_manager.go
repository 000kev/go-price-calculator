package file_manager

import (
	// "errors"
	"fmt"
	"bufio"
	"os"
	"strconv"
)

type FileManager struct {
	inputFile string
	outputFile string 
}

func New (inputFileName, outputFileName string) *FileManager {

	return &FileManager{
		inputFile: inputFileName,
		outputFile: outputFileName,
	}
}

func (fm FileManager) Load() []float64 {
	file, err := os.Open(fm.inputFile)
	prices := make([]float64, 0)
	
	if err != nil {
		fmt.Println("File does not exist")
		return nil
	}

	reader := bufio.NewScanner(file)

	for reader.Scan() {
		val, _ := strconv.ParseFloat(reader.Text(), 64)
		prices = append(prices, val)
	}

	file.Close()
	return prices
}
