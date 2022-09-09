package main

import (
	"fmt"
	"strings"

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
		if strings.Contains(item.Name, "Processor") {
			fmt.Println("Name:        ", item.Name)
			fmt.Println("HelpText:    ", item.HelpText)
			fmt.Println("CounterDefs: ")
			for _, instance := range item.Instances {
				fmt.Println("  Instance Name: ", instance.Name)
				for _, counter := range instance.Counters {
					fmt.Println("    Counter Value: ", counter.Value)
					fmt.Println("    Counter Def Name: ", counter.Def.Name)
				}
			}
			fmt.Println()
		}
	}
}
