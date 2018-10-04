package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

var problems []Problem
var result Result

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

	ask()
}

func ask() {
	go timer(3)

	reader := bufio.NewReader(os.Stdin)

	for _, p := range problems {
		fmt.Printf("Please enter the answer for this question: %v \n", p.Question)
		input, _ := reader.ReadString('\n')
		err, userAnswer := strToInt(strings.Trim(input, "\n"))
		_, standardAnswer := strToInt(p.Answer)
		correct := userAnswer == standardAnswer
		if correct && err == nil {
			result.Correct = result.Correct + 1
		} else {
			result.Incorrect = result.Incorrect + 1
		}
	}

	printResult()
}

func strToInt(s string) (error, int) {
	var answer int
	var err error
	answer, err = strconv.Atoi(s)
	return err, answer
}
