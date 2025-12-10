# Math Skills - Statistical Analysis Tool

A Go program that calculates various statistical measures from a dataset provided in a text file.

## Features

The program calculates the following statistical measures:

- **Average (Mean)** - The arithmetic mean of all values
- **Median** - The middle value when data is sorted
- **Variance** - Measures data spread using sample variance formula (n-1)
- **Standard Deviation** - Square root of variance, shows data dispersion

## Requirements

- Go 1.16 or higher

## Installation

1. Clone or download this repository
2. Navigate to the project directory

```bash
cd math-skills
```

## Usage

### Running the Program

```bash
go run . <filename>
```

**Example:**
```bash
go run . result.txt
```

### Input File Format

The input file should contain one integer per line:

```
100
152
123
123
123
165
564
100
152
123
```

### Sample Output

```
Avarage: 172.500000
Median: 123.000000
Variance: 19386.944444
Standard Deviation: 139.237449
```

## Running Tests

The project includes comprehensive unit tests for all statistical functions.

```bash
# Run all tests
go test

# Run with verbose output
go test -v

# Run specific test
go test -run TestMean

# Check test coverage
go test -cover
```

## Functions

### Main Statistical Functions

- `Mean(data []string) float64` - Calculates the arithmetic mean
- `Median(data []string) float64` - Finds the median value
- `Variance(data []string) float64` - Calculates sample variance
- `StandardDeviation(data []string) float64` - Calculates standard deviation

### Utility Functions

- `MedianInt(data []int) int` - Helper function for calculating median of integers
- `Range(data []string) int` - Calculates the range (max - min)
- `InterquartileRange(data []string) int` - Calculates IQR (Q3 - Q1)
- `Min(data []string) int` - Finds minimum value
- `Max(data []string) int` - Finds maximum value

## Implementation Notes

- **Sample Variance**: The variance function uses the sample variance formula (dividing by n-1) rather than population variance (dividing by n). This provides an unbiased estimate when working with sample data.

- **Error Handling**: The program includes basic error handling for file operations and data conversion.

## Example

Given the dataset:
```
100, 152, 123, 123, 123, 165, 564, 100, 152, 123
```

Results:
- **Mean**: 172.5 (sum of all values divided by count)
- **Median**: 123.0 (middle value in sorted data)
- **Variance**: 19386.94 (measure of spread)
- **Standard Deviation**: 139.24 (square root of variance)

## Error Messages

- `"Usage: go run yourProgram.go filename"` - No filename provided
- `"Error Opening File: ..."` - File not found or cannot be opened
- `"Cant convert those data to integers"` - Non-numeric data in file

## License

This project is available for educational purposes.

## Author

Created as a statistical analysis tool for processing numerical datasets.