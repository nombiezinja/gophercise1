package main

import (
	"fmt"
	"os"
	"time"
)

func timer(s int) {
	time.Sleep(time.Second * 3)
	fmt.Println("Time is up, exiting program.")
	printResult()
	os.Exit(0)
}
