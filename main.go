package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run yourProgram.go filename")
		return
	}
	fileName := os.Args[1]

	file, err := os.Open(fileName)
	if err != nil {
		// fmt.Printf("Error Opening File: %v", err)
		// os.Exit(1)
		log.Fatalf("Error Opening File: %v", err)
	}

	defer file.Close()
	var data []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		data = append(data, line)
	}

	fmt.Printf("Avarage: %d \n", int(math.Round(Mean(data))))
	fmt.Printf("Median: %d \n", int(math.Round(Median(data))))
	fmt.Printf("Range: %d \n", Range(data))
	fmt.Printf("Interquartile: %d \n", InterquartileRange(data))
	fmt.Printf("Min Value: %d \n", Min(data))
	fmt.Printf("Max Value: %d \n", Max(data))
	fmt.Printf("Variance: %d \n", int(math.Round(Variance(data))))
	fmt.Printf("Standard Deviation: %d \n", int(math.Round(StandardDeviation(data))))

}

func Mean(data []string) float64 {
	if len(data) == 0 {
		return 0
	}
	var mean float64
	sum := 0
	n := len(data)
	for _, num := range data {
		numConv, err := strconv.Atoi(num)
		if err != nil {
			fmt.Printf("Cant convert those data to integers")
		}
		sum += numConv
	}
	mean = float64(sum) / float64(n)
	return mean
}

func Median(data []string) float64 {
	var median float64
	var dataConv []float64
	for _, num := range data {
		numConv, err := strconv.Atoi(num)
		if err != nil {
			fmt.Printf("Cant convert those data to integers")
		}
		dataConv = append(dataConv, float64(numConv))
	}
	sort.Float64s(dataConv)
	if (len(dataConv))%2 == 0 {
		first := dataConv[((len(dataConv))/2)-1]
		second := dataConv[((len(dataConv)) / 2)]
		median = ((first + second) / 2)
	} else {
		median = dataConv[(len(dataConv))/2]
	}
	return median
}

func MedianInt(data []int) int {
	var median int
	sort.Ints(data)
	if (len(data))%2 == 0 {
		first := data[((len(data))/2)-1]
		second := data[((len(data)) / 2)]
		median = ((first + second) / 2)
	} else {
		median = data[(len(data))/2]
	}
	return median
}

func Range(data []string) int {
	var Range int
	var dataConv []int
	for _, num := range data {
		numConv, err := strconv.Atoi(num)
		if err != nil {
			fmt.Printf("Cant convert those data to integers")
		}
		dataConv = append(dataConv, numConv)
	}
	sort.Ints(dataConv)
	maxVal := dataConv[len(dataConv)-1]
	minVal := dataConv[0]
	Range = maxVal - minVal
	return Range
}

func InterquartileRange(data []string) int {
	var iqr int
	var dataConv []int
	var q1 int
	var q3 int
	n := len(data)
	half := n / 2
	for _, num := range data {
		numConv, err := strconv.Atoi(num)
		if err != nil {
			fmt.Printf("Cant convert those data to integers")
		}
		dataConv = append(dataConv, numConv)
	}
	sort.Ints(dataConv)
	if n%2 == 0 {
		q1 = MedianInt(dataConv[:half])
		q3 = MedianInt(dataConv[half:])
	} else {
		q1 = MedianInt(dataConv[:half])
		q3 = MedianInt(dataConv[half+1:])
	}
	iqr = q3 - q1
	return iqr
}

func Min(data []string) int {
	var dataConv []int
	for _, num := range data {
		numConv, err := strconv.Atoi(num)
		if err != nil {
			fmt.Printf("Cant convert those data to integers")
		}
		dataConv = append(dataConv, numConv)
	}
	sort.Ints(dataConv)
	return dataConv[0]
}

func Max(data []string) int {
	var dataConv []int
	n := len(data)
	for _, num := range data {
		numConv, err := strconv.Atoi(num)
		if err != nil {
			fmt.Printf("Cant convert those data to integers")
		}
		dataConv = append(dataConv, numConv)
	}
	sort.Ints(dataConv)
	return dataConv[n-1]
}

func Variance(data []string) float64 {
	var variance float64
	var total float64
	n := len(data)
	mean := Mean(data)
	for _, num := range data {
		numConv, err := strconv.Atoi(num)
		if err != nil {
			fmt.Println("Can not Convert to number")
		}
		diff := math.Pow((mean - float64(numConv)), 2)
		total += diff
	}
	variance = total / float64(n)
	return variance
}

func StandardDeviation(data []string) float64 {
	return math.Sqrt(Variance(data))
}
