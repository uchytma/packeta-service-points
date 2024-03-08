package main

import (
	"fmt"
	"packeta"
)

func main() {
	// downloadData()

	ch := make(chan packeta.CarrierPoint)
	go packeta.Consumer(ch)

	err := packeta.ProcessFile(ch)
	if err != nil {
		fmt.Println(err)
	}
}
