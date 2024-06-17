package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"log"
	"math"
	"os"
	"slices"
	"strconv"
)

type dataframe struct {
	headers []string
	rows    [][]string
}

func (df *dataframe) getHeaders() []string {
	return df.headers
}

func (df *dataframe) getRow(rowIndex int) ([]string, error) {
	if rowIndex < 0 || rowIndex >= len(df.rows) {
		return nil, errors.New("index out of bounds")
	}

	return df.rows[rowIndex], nil
}

func (df *dataframe) getColumn(columnIndex int) ([]string, error) {
	if columnIndex < 0 || columnIndex >= len(df.headers) {
		return nil, errors.New("index out of bounds")
	}

	var newColumn []string = make([]string, len(df.rows))
	for _, row := range df.rows {
		newColumn = append(newColumn, row[columnIndex])
	}

	return newColumn, nil
}

// TODO: handle missing data somehow
func convertToFloat(series []string) ([]float64, error) {
	var floatValues []float64 = make([]float64, 0)

	for _, strVal := range series {
		floatValue, err := strconv.ParseFloat(strVal, 64)
		if err != nil {
			return nil, err
		}
		floatValues = append(floatValues, floatValue)
	}

	return floatValues, nil
}

func average(nums []float64) float64 {
	total := 0.0
	for _, num := range nums {
		total += num
	}
	return total / float64(len(nums))
}

func median(nums []float64) float64 {
	n := len(nums)
	sorted := make([]float64, n)
	copy(sorted, nums)
	slices.Sort(sorted)

	if n%2 == 0 {
		return (sorted[n/2-1] + sorted[n/2]) / 2
	}
	return sorted[n/2]
}

func standardDeviation(nums []float64) float64 {
	mean := average(nums)
	var sumSquares float64
	for _, num := range nums {
		sumSquares += (num - mean) * (num - mean)
	}
	return math.Sqrt(sumSquares / float64(len(nums)))
}

func loadcsv(filepath string) (*dataframe, error) {
	var df dataframe = dataframe{}
	f, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	r := csv.NewReader(f)
	rows, err := r.ReadAll()
	if err != nil {
		return nil, err
	}

	if len(rows) <= 1 {
		df.rows = rows
		return &df, nil
	}
	df.headers = rows[0]

	df.rows = rows[1:]

	return &df, nil
}

func main() {
	// open file
	// read contents into slice of slices
	// make df with data
	// start loop
	// choose command and row or column
	// get row or column
	// convert to floats
	// run stat function
	// restart loop

	df, err := loadcsv(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	for {
		var userCommand string
		_, err := fmt.Scanln(&userCommand)
		if err != nil {
			continue
		}

		switch userCommand {
		case "avg":
			for {
				var userRow string
				var userColumn string
				_, err = fmt.Scanln(&userRow)
				_, err = fmt.Scanln(&userColumn)

				if userRow != "" {

				}

			}
		case "stdv":
		case "median":
		case "q":
		default:
		}
	}

}
