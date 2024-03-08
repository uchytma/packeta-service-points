package main

import "fmt"

func main() {
	// downloadData()

	ch := make(chan CarrierPoint)
	go consumer(ch)

	err := processFile(ch)
	if err != nil {
		fmt.Println(err)
	}
}
