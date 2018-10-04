package main

import "fmt"

func printResult() {

	fmt.Printf(
		"You've answered %v questions correctly, and %v questions incorrectly \n",
		result.Correct,
		result.Incorrect,
	)
}
