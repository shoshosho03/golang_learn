package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
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

	//context(終了制御用)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	//Ctrl+Cのシグナルを受信
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	//引数で受け取った値でtickerを作成
	ticker := time.NewTicker(*interval)
	defer ticker.Stop()

	//シグナル待ちgoroutine
	go func() {
		sig := <-sigChan
		fmt.Println("\nreceived", sig)
		cancel()
	}()

	fmt.Println("start monioring")

	for {
		select {
		case <-ctx.Done():
			fmt.Println("graceful shutdown")
			return

		case <-ticker.C:
			memPercent, err := mem.VirtualMemory()
			if err != nil {
				fmt.Println("mem error", err)
				continue
			}

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
}
