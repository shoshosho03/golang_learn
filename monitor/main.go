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
	//引数で受け取った値でtickerを作成
	ticker := time.NewTicker(*interval)
	defer ticker.Stop()

	for range ticker.C {
		// メモリ使用率取得
		memPercent, err := mem.VirtualMemory()
		if err != nil {
			fmt.Println("mem error", err)
			continue
		}
		// CPU使用率取得
		cpuPercent, err := cpu.Percent(0, false)
		if err != nil {
			fmt.Println("cpu error", err)
			continue

		}
		// 出力
		fmt.Printf("UsedPercent:%.2f%%\n", memPercent.UsedPercent)
		fmt.Printf("CPU Percent: %.2f%%\n", cpuPercent[0])
		fmt.Println("----")

	}
}
