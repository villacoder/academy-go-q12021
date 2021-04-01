package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type csvLine struct {
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Publisher string `json:"publisher"`
	Place     string `json:"place_of_publication"`
	Year      int    `json:"start_year"`
}

func ReadCsv(filename string) ([][]string, error) {

	// Open CSV file
	f, err := os.Open(filename)
	if err != nil {
		return [][]string{}, err
	}
	defer f.Close()

	// Read File into a Variable
	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return [][]string{}, err
	}

	return lines, nil
}

func csvToMap() (newspapers []csvLine) {
	var rows []csvLine
	lines, err := ReadCsv("data/newspapers.csv")

	if err != nil {
		fmt.Println(err)
	}

	for _, line := range lines {
		id, _ := strconv.Atoi(line[0])
		year, _ := strconv.Atoi(line[4])

		data := csvLine{
			Id:        id,
			Title:     line[1],
			Publisher: line[2],
			Place:     line[3],
			Year:      year,
		}
		rows = append(rows, data)

	}

	return rows
}
