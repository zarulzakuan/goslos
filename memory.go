package main

import (
	"runtime"
	"time"

	memutil "github.com/shirou/gopsutil/mem"
)

func RunMemLoad(maxMemPercentage uint) {
	operationInterval := 500
	stackSize := 1 * 1024 * 1024 * 1024
	go func() {
		dummy := [][]byte{}
		totalStacked := 0
		for {
			stack := make([]byte, stackSize)
			totalStacked += len(stack)
			dummy = append(dummy, stack)
			time.Sleep(time.Duration(operationInterval) * time.Millisecond)
			vmstat, _ := memutil.VirtualMemory()
			if vmstat.UsedPercent > float64(maxMemPercentage) {
				// run GC to clear unused memory
				runtime.GC()
				break
			}
		}
	}()

}
