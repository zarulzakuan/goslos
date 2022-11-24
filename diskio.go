package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func RunDiskIO(minChunkSizeInByte string, maxChunkSizeInByte string, totalSizeInByte string, operationIntervalinMilisecond int, writeRatioInPercentage float32) {
	//writtenByte := []byte{}

	minChunkSizeInByteParsed := ParseUnitToByte(minChunkSizeInByte)
	maxChunkSizeInByteParsed := ParseUnitToByte(maxChunkSizeInByte)
	totalSizeInByteParsed := ParseUnitToByte(totalSizeInByte)

	f, err := os.Create(fmt.Sprintf("%s/golsim", os.TempDir()))
	defer func() {
		f.Close()
		os.Remove(fmt.Sprintf("%s/golsim", os.TempDir()))
	}()
	if err != nil {
		println(err)
		return
	}

	var writeCount float32 = 0
	var readCount float32 = 0
	var currentWriteRatio float32 = 0
	var totalChunkWritten int64 = 0

start:
	for int64(totalChunkWritten) < totalSizeInByteParsed {

		chunkSize := GetChunk(minChunkSizeInByteParsed, maxChunkSizeInByteParsed)
		totalChunkToBeWritten := totalChunkWritten + chunkSize
		// println("Total written: ", totalChunkWritten)

		// check how many bytes we are going write, if exceed, break
		if totalChunkToBeWritten > totalSizeInByteParsed {

			println("All written")
			break
		}

		chunk := make([]byte, chunkSize)

		// intially we write and read
		if writeCount == 0 && readCount == 0 {
			totalChunkToBeWritten = totalChunkToBeWritten + chunkSize
			// println("Total to bewritten: ", totalChunkToBeWritten)
			Write(f, chunk, &totalChunkWritten)
			writeCount++
			Read(f, chunk)
			readCount++
		} else {
			// println("Total to bewritten: ", totalChunkToBeWritten)
		}

		currentWriteRatio = 100.0 - (readCount / writeCount * 100.0)

		if currentWriteRatio > writeRatioInPercentage {
			Write(f, chunk, &totalChunkWritten)
			writeCount++
			Read(f, chunk)
			readCount++
			currentWriteRatio = 0
		} else {
			Write(f, chunk, &totalChunkWritten)
			writeCount++
		}

		time.Sleep(time.Duration(operationIntervalinMilisecond) * time.Millisecond)
	}

	// finished writing, so we truncate the file and write/read again from truncated position
	println("Truncate!")
	truncatedPosition := GetChunk(minChunkSizeInByteParsed, totalSizeInByteParsed)
	f.Truncate(truncatedPosition)
	f.Seek(truncatedPosition, 0)
	totalChunkWritten -= truncatedPosition
	if totalChunkWritten < 0 {
		totalChunkWritten = 0
	}
	println("Total written after truncate: ", totalChunkWritten)
	goto start
}

func Write(f *os.File, bytes []byte, lastWritePos *int64) {
	f.Seek(*lastWritePos, 0)
	n, err := f.Write(bytes)
	// _ = n
	bytes = nil
	if err != nil {
		println(err)
		return
	}
	// println(n, "bytes written <----")
	*lastWritePos += int64(n)
	if err != nil {
		println(err)
		return
	}
}

func Read(f *os.File, bytes []byte) {
	_, err := f.Seek(0, 0)

	if err != nil {
		println(err)
		return
	}
	n, err := f.Read(bytes)
	_ = n
	// println(n, " bytes read ---->")
	bytes = nil
	if err != nil {
		println(err)
		return
	}

}

func GetChunk(min int64, max int64) int64 {
	rand.Seed(time.Now().UnixNano())
	return rand.Int63n(max-min+1) + min

}
