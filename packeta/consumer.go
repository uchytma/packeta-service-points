package packeta

import (
	"fmt"
)

// Consumer simulates a function that consumes data received on the provided channel.
func Consumer(ch <-chan CarrierPoint) {
	for point := range ch {
		fmt.Println(point.String())
	}
}
