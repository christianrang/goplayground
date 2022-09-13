package main

// Queries
// Get all Processes: Get-CimInstance -Query 'Select * from Win32_Process'

import (
	"fmt"
	"os"

	"github.com/StackExchange/wmi"
)

const win32ProcessorQuery = "SELECT LoadPercentage FROM Win32_Processor"

type win32Processor struct {
	LoadPercentage uint32
}

func main() {
	var dst []win32Processor

	if err := wmi.Query(win32ProcessorQuery, &dst); err != nil {
		fmt.Println("encountered error querying WMI", err)
		os.Exit(1)
	}

	fmt.Printf("%+v\n", dst)
}
