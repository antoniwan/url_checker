# Ping URLs from a CSV File

This Go program scans the script's folder for `.csv` files, lets you select one of the CSV files based on user input, and then processes the CSV file to:

1. Validate if the first column of each row is a valid URL.
2. If valid, pings the URL to get its HTTP status code.
3. Adds two new columns to the CSV:
   - `url_validity`: Indicates whether the URL is valid.
   - `status_code`: The HTTP status code of the URL.

## Requirements

- Go (https://golang.org/doc/install)

## Installation

1. Clone the repository or download the `csv_updater.go` file.
2. Ensure you have Go installed on your system.

## Usage

1. Place the `csv_updater.go` file in a directory containing the CSV files you want to process.
2. Open a terminal and navigate to the directory containing `csv_updater.go`.
3. Run the program using the following command:

   ```sh
   go run csv_updater.go
   ```

4. The program will list all `.csv` files in the directory. Select a file by entering its corresponding number.
5. The program will read the selected CSV file, validate the URLs, get their status codes, and update the CSV file with the new columns.

## Example

I've included 2 sample CSV files to the project
