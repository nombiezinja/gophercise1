package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

type Problem struct {
	problem string
	answer  string
}

func main() {
	csvFile, _ := os.Open("problems.csv")
	reader := csv.NewReader(bufio.NewReader(csvFile))
	var problem []Problem
	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		problem = append(problem, Problem{
			problem: line[0],
			answer:  line[1],
		})
	}
	// quiz, _ := problem
	fmt.Println(problem)
	fmt.Println("hello")
}
