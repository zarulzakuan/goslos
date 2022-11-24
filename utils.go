package main

import (
	"log"
	"strconv"
	"strings"
)

func ParseUnitToByte(s string) int64 {
	var l, n []rune
	var unitInByte int64 = 0
	for _, r := range s {
		switch {
		case r >= 'A' && r <= 'Z':
			l = append(l, r)
		case r >= 'a' && r <= 'z':
			l = append(l, r)
		case r >= '0' && r <= '9':
			n = append(n, r)
		}
	}

	switch strings.ToUpper(string(l)) {
	case "G":
		unitInByte = convertStringToInt64(string(n)) * 1024 * 1024 * 1024
	case "M":
		unitInByte = convertStringToInt64(string(n)) * 1024 * 1024
	case "K":
		unitInByte = convertStringToInt64(string(n)) * 1024
	case "B":
		unitInByte = convertStringToInt64(string(n))
	default:
		log.Fatalln(`Invalid unit. Accepting "G", "M", "K", "B" `)
	}

	return unitInByte
}

func convertStringToInt64(s string) int64 {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		panic(err)
	}
	return i
}
