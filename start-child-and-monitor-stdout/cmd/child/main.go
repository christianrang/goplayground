package main

import (
	"fmt"
	"time"
)

func main() {
	values := []string{"2023/07/11 asdf 1", "2023/07/10 foo 2", "2023/07/09 bar 3"}
	for {
		for _, value := range values {
			fmt.Println(value)
			time.Sleep(time.Second)
		}
	}
}
