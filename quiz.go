package main

import (
	"bufio"
	"encoding/csv"
	"io"
	"log"
	"os"
)

// parse csv into slice of struct containing question and answer
// range through slice and prompt for answer to each question
// increment correct/incorrect accordingly
// at the end of the slice, display the results

var problems []Problem

func main() {
	csvFile, _ := os.Open("problems.csv")
	reader := csv.NewReader(bufio.NewReader(csvFile))

	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		problems = append(problems, Problem{
			Question: line[0],
			Answer:   line[1],
		})
	}
}

func ask() {
	for q := range problems {

	}
}
