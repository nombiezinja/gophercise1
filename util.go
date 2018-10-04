package main

import (
	"os"
	"strconv"
)

func parseTimeLimit() int {
	t, err := strconv.Atoi(os.Args[1])
	FailOnError(err, "Failed to parse commandline argument for time limit")
	return t
}

func strToInt(s string) (error, int) {
	var a int
	var err error
	a, err = strconv.Atoi(s)
	return err, a
}
