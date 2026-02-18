package main

import (
	"fmt"
	"time"

	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/mem"
)

func main() {
	v, _ := mem.VirtualMemory()
	c, _ := cpu.Percent(time.Second, false)
	// almost every return value is a struct
	//fmt.Printf("TotalMem: %d, Free:%d, UsedPercent:%.2f%%\n", v.Total, v.Free, v.UsedPercent)
	fmt.Printf("UsedPercent:%.2f%%\n", v.UsedPercent)
	fmt.Printf("CPU Percent: %.2f%%\n", c[0])

	// convert to JSON. String() is also implemented
	//fmt.Println(v)
	//fmt.Println(c)

	fmt.Println("Monitoring system resources for 3 seconds...")
	time.Sleep(3 * time.Second)
}
