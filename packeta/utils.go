package packeta

import (
	"fmt"
	"os"
)

func getPacketaUrl() string {
	packetaKey := os.Getenv("PACKETA_KEY")
	return fmt.Sprintf("https://www.zasilkovna.cz/api/%s/point.json", packetaKey)
}

const filePath = "files/output.json"
