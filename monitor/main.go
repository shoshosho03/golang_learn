package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/mem"
)

func main() {
	// 監視間隔をコマンドライン引数から取得
	interval := flag.Duration("interval", 1*time.Second, "監視間隔を秒単位で指定")
	flag.Parse()

	if *interval <= 0 {
		fmt.Println("Interval must be greater than 0")
		return
	}

	v, err := mem.VirtualMemory()
	if err != nil {
		fmt.Println("Error getting virtual memory:", err)
		return
	}

	c, err := cpu.Percent(time.Second, false)
	if err != nil {
		fmt.Println("Error getting CPU percent:", err)
		return
	}

	for {

		// システムリソースの使用状況を表示
		fmt.Printf("UsedPercent:%.2f%%\n", v.UsedPercent)
		fmt.Printf("CPU Percent: %.2f%%\n", c[0])

		// 指定された間隔だけ待機
		fmt.Println("Monitoring system resources for", *interval)
		time.Sleep(*interval)
	}

}
