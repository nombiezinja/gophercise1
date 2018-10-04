package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ask() {
	go timer(3)

	result.Unanswered = len(problems)
	reader := bufio.NewReader(os.Stdin)

	for _, p := range problems {
		fmt.Printf("Please enter the answer for this question: %v \n", p.Question)
		input, _ := reader.ReadString('\n')
		err, userAnswer := strToInt(strings.Trim(input, "\n"))
		_, standardAnswer := strToInt(p.Answer)
		correct := userAnswer == standardAnswer

		if err == nil {
			result.Unanswered = result.Unanswered - 1
		}

		if correct {
			result.Correct = result.Correct + 1
		} else {
			result.Incorrect = result.Incorrect + 1
		}
	}

	printResult()
}
