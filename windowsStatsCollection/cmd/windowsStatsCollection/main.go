package main

import (
	"fmt"

	"github.com/leoluk/perflib_exporter/perflib"
	// "os"
	// "syscall"
	// "unsafe"
)

func main() {
	// var (
	// 	bufferSize uint32
	// 	valType    uint32
	// )
	// name, err := syscall.UTF16PtrFromString("Global")
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }
	// bufferSize = 4096
	// buffer := make([]byte, bufferSize)
	// defer syscall.RegCloseKey(syscall.HKEY_PERFORMANCE_DATA)
	// err = syscall.RegQueryValueEx(
	// 	syscall.HKEY_PERFORMANCE_DATA,
	// 	name,
	// 	nil,
	// 	&valType,
	// 	(*byte)(unsafe.Pointer(&buffer[0])),
	// 	&bufferSize,
	// )
	// fmt.Println(string(buffer))

	obj, _ := perflib.QueryPerformanceData("Global")
	for _, item := range obj {
		// if item.Name == "Processor" || item.Name == "Processor Information" {
		if item.Name == "Processor Information" {
			// fmt.Println("Name:        ", item.Name)
			// fmt.Println("HelpText:    ", item.HelpText)
			for _, instance := range item.Instances {
				var (
					total int64
					user  int64
				)
				// fmt.Println("  Instance Name: ", instance.Name)
				for _, counter := range instance.Counters {
					// fmt.Println("    Counter Def Name: ", counter.Def.Name)
					// fmt.Println("    Counter Value: ", counter.Value)
					if counter.Def.Name == "% User Time" {
						user = counter.Value
					}
					if counter.Def.Name == "% Processor Time" {
						total = counter.Value
					}

				}
				if instance.Name == "_Total" && total > 0 {
					fmt.Println("User:      ", user)
					fmt.Println("Total:     ", total)
					fmt.Println("Perc user: ", float64(float64(user)/float64(total)))
				}
			}
		}
	}
}
