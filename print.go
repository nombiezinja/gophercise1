package main

import "fmt"

func printResult() {

	fmt.Printf(
		"You've answered %v questions correctly, %v questions incorrectly. You have left %v questions unanswered. \n",
		result.Correct,
		result.Incorrect,
		result.Unanswered,
	)
}
