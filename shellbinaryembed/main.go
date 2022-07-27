package main

import (
	_ "embed"
	"fmt"
	"os/exec"
	"strings"
)

//go:embed echo.sh
var echoSh string

func main() {
	fmt.Println(echoSh)
	c := exec.Command("bash")
	c.Stdin = strings.NewReader(echoSh)

	b, e := c.Output()
	if e != nil {
		fmt.Println(e)
	}
	fmt.Println(string(b))
}
