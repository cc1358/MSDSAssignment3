package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/spf13/cobra"
)

// Declare variables
var inputFilePath string
var outputFilePath string
var iterations int

// Set up command line interface with Cobra
func main() {
	var rootCmd = &cobra.Command{Use: "housing-stats", Run: run}

	rootCmd.Flags().StringVarP(&inputFilePath, "input", "i", "housesInput.csv", "Input CSV file path")
	rootCmd.Flags().StringVarP(&outputFilePath, "output", "o", "housesOutputGo.txt", "Output text file path")
	rootCmd.Flags().IntVarP(&iterations, "iterations", "n", 100, "Number of iterations")

}

// Execute housing-stats command
func run(cmd *cobra.Command, args []string) {
	startTime := time.Now()

	printHeader()

	for i := 0; i < iterations; i++ {
		processData()
	}

	elapsedTime := time.Since(startTime)
	fmt.Printf("CPU Processing Time: %v\n", elapsedTime)
}

// Organize header information in CSV file to format output
func printHeader() {
	file, err := os.Open(inputFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	headers, err := reader.Read()
	if err != nil {
		log.Fatal(err)
	}

	headerLine := fmt.Sprintf("%-15s%-15s%-15s%-15s%-15s%-15s%-15s\n", headers[0], headers[1], headers[2], headers[3], headers[4], headers[5], headers[6])
	fmt.Print(headerLine)

	outputFile, err := os.OpenFile(outputFilePath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer outputFile.Close()

	outputFile.WriteString(headerLine)
}

// Process data from input file
func processData() {
	file, err := os.Open(inputFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	outputFile, err := os.OpenFile(outputFilePath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer outputFile.Close()

	headers := records[0]

	headerLine := fmt.Sprintf("%-15s%-15s%-15s%-15s%-15s%-15s%-15s\n", headers[0], headers[1], headers[2], headers[3], headers[4], headers[5], headers[6])
	fmt.Print(headerLine)
	outputFile.WriteString(headerLine)

	var minValues, maxValues, meanValues []float64

	// Calculate summary statistics for each column
	for colIndex := 0; colIndex < len(records[0]); colIndex++ {
		var columnValues []float64
		for rowIndex := 1; rowIndex < len(records); rowIndex++ {
			value, err := strconv.ParseFloat(records[rowIndex][colIndex], 64)
			if err != nil {
				log.Printf("Error converting value to float: %v", err)
				// Skip the current iteration if there's an error
				continue
			}
			columnValues = append(columnValues, value)
		}

		// Calculate statistics for the column
		min, max, mean := calculateStatistics(columnValues)
		minValues = append(minValues, min)
		maxValues = append(maxValues, max)
		meanValues = append(meanValues, mean)

		// Write the result to the output file
		result := fmt.Sprintf("%-15f%-15f%-15f", min, max, mean)
		outputFile.WriteString(result)
	}

	// Add newline after writing all columns' statistics
	outputFile.WriteString("\n")
}

// Caclulate minimum, maximm and mean values
func calculateStatistics(values []float64) (float64, float64, float64) {
	var sum, min, max float64
	min = 1<<63 - 1
	max = -1 << 63

	for _, value := range values {
		sum += value

		if value < min {
			min = value
		}

		if value > max {
			max = value
		}
	}

	// Check if the values contain valid numeric values
	if len(values) == 0 {
		log.Println("No valid numeric values found.")
		return 0, 0, 0
	}

	mean := sum / float64(len(values))
	return min, max, mean
}
