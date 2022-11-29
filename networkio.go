package main

import (
	"fmt"
	"net/http"
	"time"
)

// call external API with random response size
// TODO - need to simulate request size as well
func RunNetworkIO() {

	baseURL := "https://pokeapi.co/api/v2/pokemon?limit="

	for {
		rand := RandomGenerator(100, 100000)
		randWait := RandomGenerator(1000, 5000)
		_, err := http.Get(fmt.Sprintf("%s%d", baseURL, rand))
		if err != nil {
			fmt.Print(err.Error())
		}

		time.Sleep(time.Duration(randWait) * time.Millisecond)
	}
}
