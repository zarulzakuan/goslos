package main

import (
	"time"
)

func RunMemLoad(timeSeconds int, memorySizeToUse string) {

	memoryParsed := ParseUnitToByte(memorySizeToUse)

	go func() {

		var original = make([]byte, memoryParsed)
		_ = original
		time.Sleep(time.Duration(timeSeconds) * time.Second)
		print("Time's up!")

	}()
}
