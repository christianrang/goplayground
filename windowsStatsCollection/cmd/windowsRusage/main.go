package main

import (
	"fmt"
	"os"
	"time"

	"golang.org/x/sys/windows"
)

func main() {
	time.Sleep(time.Second * 5)

	values, err := GetRUsage()
	if err != nil {
		fmt.Println("error: ", err)
		os.Exit(1)
	}

	fmt.Printf("%+v\n", values)
}

func GetRUsage() (windows.Rusage, error) {
	var values windows.Rusage

	currentProcessHandler, err := windows.GetCurrentProcess()
	if err != nil {

		return values, err
	}

	windows.GetProcessTimes(currentProcessHandler, &values.CreationTime, &values.ExitTime, &values.KernelTime, &values.UserTime)

	return values, nil
}
