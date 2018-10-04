package main

import (
	"os"
	"strconv"
	"strings"
)

func parseTimeLimit() int {
	if len(os.Args) == 2 {
		t, err := strconv.Atoi(os.Args[1])
		FailOnError(err, "Please pass an integer for time limit.")
		return t
	} else {
		return 30
	}
}

func strToInt(s string) (error, int) {
	var a int
	var err error
	a, err = strconv.Atoi(s)
	return err, a
}

func normalize(s string) string {
	return strings.Join(strings.Fields(s), " ")
}
