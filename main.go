package main

import (
	"fmt"
	"time"
)

func main() {
	for {
		fmt.Println("Hello World! This is Go.")
		time.Sleep(5 * time.Minute)
	}
}
