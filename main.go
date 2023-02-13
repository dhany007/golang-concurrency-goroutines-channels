package main

import (
	"fmt"
	"time"
)

func timesThree(number int) {
	fmt.Println(number * 3)
}

func main() {
	fmt.Println("we are executing a goroutine")
	go timesThree(2)
	fmt.Println("done")
	time.Sleep(time.Second) // prevents the program exiting before the goroutine start
}

// output:
// we are executing a goroutine
// done
// 6
