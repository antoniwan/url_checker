package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func isValidURL(testURL string) bool {
	_, err := url.ParseRequestURI(testURL)
	return err == nil
}

func getStatusCode(testURL string) int {
	client := http.Client{
		Timeout: 5 * time.Second,
	}
	resp, err := client.Get(testURL)
	if err != nil {
		return 0 // Return 0 if there's an error
	}
	defer resp.Body.Close()
	return resp.StatusCode
}

func main() {
	// Get the current directory
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current directory:", err)
		return
	}

	// List all .csv files in the current directory
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	var csvFiles []string
	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".csv") {
			csvFiles = append(csvFiles, file.Name())
		}
	}

	if len(csvFiles) == 0 {
		fmt.Println("No .csv files found in the current directory.")
		return
	}

	// Display the list of CSV files
	fmt.Println("Select a CSV file to read:")
	for i, file := range csvFiles {
		fmt.Printf("%d: %s\n", i+1, file)
	}

	// Get the user's choice
	var choice int
	fmt.Print("Enter the number of the file you want to read: ")
	_, err = fmt.Scanf("%d", &choice)
	if err != nil || choice < 1 || choice > len(csvFiles) {
		fmt.Println("Invalid choice.")
		return
	}

	selectedFile := csvFiles[choice-1]
	filePath := filepath.Join(dir, selectedFile)

	// Open the selected CSV file
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Read the CSV file
	reader := csv.NewReader(bufio.NewReader(file))

	// Read the header
	header, err := reader.Read()
	if err != nil {
		fmt.Println("Error reading CSV header:", err)
		return
	}

	// Add the new "url_validity" and "status_code" columns to the header
	header = append(header, "url_validity", "status_code")

	// Create a slice to hold the data rows
	var rows [][]string

	// Read the rest of the rows
	for {
		record, err := reader.Read()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			fmt.Println("Error reading CSV file:", err)
			return
		}

		// Check if the first column is a valid URL
		urlValidity := "read"
		statusCode := "0"
		if isValidURL(record[0]) {
			urlValidity = "valid_url"
			statusCode = fmt.Sprintf("%d", getStatusCode(record[0]))
		}

		// Add the "url_validity" and "status_code" columns to the record
		record = append(record, urlValidity, statusCode)
		rows = append(rows, record)
	}

	// Display the header and rows
	fmt.Println("Header:", header)
	fmt.Println("Rows:")
	for i, row := range rows {
		fmt.Printf("Row %d: %v\n", i+1, row)
	}

	// Write the updated data back to the CSV file
	outputFile, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer outputFile.Close()

	writer := csv.NewWriter(outputFile)
	defer writer.Flush()

	// Write the header
	if err := writer.Write(header); err != nil {
		fmt.Println("Error writing header to CSV file:", err)
		return
	}

	// Write the rows
	for _, row := range rows {
		if err := writer.Write(row); err != nil {
			fmt.Println("Error writing row to CSV file:", err)
			return
		}
	}

	fmt.Println("CSV file updated successfully.")
}