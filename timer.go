package main

import (
	"fmt"
	"os"
	"time"
)

func timer(s int) {
	time.Sleep(time.Second * time.Duration(s))
	fmt.Println("Time is up, exiting program.")
	printResult()
	os.Exit(0)
}
