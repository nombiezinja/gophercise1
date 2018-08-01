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

// step 1
// parse csv into slice of struct containing question and answer
// range through slice and prompt for answer to each question
// increment correct/incorrect accordingly
// at the end of the slice, display the results

// step 2
// start a timer at the beginning of prompt
// display result when either timer runs out or when all q's are answered

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

	ask()
}

func ask() {
	var result Result

	reader := bufio.NewReader(os.Stdin)

	for _, p := range problems {
		fmt.Printf("Please enter the answer for this question: %v \n", p.Question)
		input, _ := reader.ReadString('\n')
		userAnswer := strToInt(strings.Trim(input, "\n"))
		standardAnswer := strToInt(p.Answer)
		correct := userAnswer == standardAnswer
		if correct {
			result.Correct = result.Correct + 1
		} else {
			result.Incorrect = result.Incorrect + 1
		}
	}

	fmt.Printf(
		"You've answered %v questions correctly, and %v questions incorrectly \n",
		result.Correct,
		result.Incorrect,
	)
}

func strToInt(s string) int {
	var answer int
	var err error
	answer, err = strconv.Atoi(s)
	if err != nil {
		FailOnError(err, "Answer entered is not a number")
	}
	return answer
}
