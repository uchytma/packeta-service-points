package packeta

import (
	"fmt"
	"testing"
)

func TestImport(t *testing.T) {
	// downloadData()

	ch := make(chan CarrierPoint)
	go Consumer(ch)

	err := ProcessFile(ch)
	if err != nil {
		fmt.Println(err)
	}
}
