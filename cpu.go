package main

import (
	"math/rand"
	"runtime"
	"time"
)

const unitHundresOfMicrosecond = 1000

func RunCPULoad(coresCount int, timeSeconds int, minUsage int, maxUsage int) {
	runtime.GOMAXPROCS(coresCount)

	// second     ,s  * 1
	// millisecond,ms * 1000
	// microsecond,μs * 1000 * 1000
	// nanosecond ,ns * 1000 * 1000 * 1000

	// every loop : run + sleep = 1 unit

	// 1 unit = 100 ms may be the best

	allQuit := []chan struct{}{}
	defaultCoreCount := runtime.NumCPU()
	if coresCount == 0 || coresCount > defaultCoreCount {
		coresCount = defaultCoreCount
	}
	if timeSeconds == 0 {
		timeSeconds = 31536000 // 1 year
	}

	for i := 0; i < coresCount; i++ {

		quit := make(chan struct{})

		allQuit = append(allQuit, quit)
		go func() {
			runtime.LockOSThread()
			// endless loop
			for {
				randCPUUsage := RandCPUUsage(minUsage, maxUsage)
				println(randCPUUsage)
				runMicrosecond := unitHundresOfMicrosecond * randCPUUsage
				sleepMicrosecond := unitHundresOfMicrosecond*100 - runMicrosecond

				select {
				case <-quit:
					return
				default:
					begin := time.Now()
					for {
						// run 100%
						if time.Since(begin) > time.Duration(runMicrosecond)*time.Microsecond { //
							break
						}
					}
				}
				// sleep
				time.Sleep(time.Duration(sleepMicrosecond) * time.Microsecond)
			}

		}()
	}
	// how long

	time.Sleep(time.Duration(timeSeconds) * time.Second)
	// print("Time's up!")

	for _, quit := range allQuit {
		quit <- struct{}{}
	}

	// print("Bye")
}

func RandCPUUsage(min int, max int) int {
	rand.Seed(time.Now().UnixNano())

	return rand.Intn(max-min+1) + min

}
