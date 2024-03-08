package packeta

import (
	"io"
	"net/http"
	"os"
)

func downloadData() error {
	out, err := os.Create(filePath)
	if err != nil {
		return err
	}
	resp, err := http.Get(getPacketaUrl())
	if err != nil {
		return err
	}

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	err = resp.Body.Close()
	if err != nil {
		return err
	}

	return out.Close()
}
