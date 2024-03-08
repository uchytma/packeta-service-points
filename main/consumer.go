package main

import "fmt"

// consumer simulates a function that consumes data received on the provided channel.
func consumer(ch <-chan CarrierPoint) {
	for point := range ch {
		fmt.Println(point.String())
	}
}
