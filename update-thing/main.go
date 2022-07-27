package main

import (
	"fmt"
	"time"
)

func main() {
	for {
		fmt.Println("Goodbye")
		time.Sleep(time.Second * 5)
	}
}
