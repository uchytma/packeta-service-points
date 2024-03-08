package packeta

import (
	"encoding/json"
	"os"
)

// skip skips the next value in the JSON document.
func skip(d *json.Decoder) error {
	n := 0
	for {
		t, err := d.Token()
		if err != nil {
			return err
		}
		switch t {
		case json.Delim('['), json.Delim('{'):
			n++
		case json.Delim(']'), json.Delim('}'):
			n--
		}
		if n == 0 {
			return nil
		}
	}
}

func processCarrierPoints(d *json.Decoder, carrier *Carrier, ch chan<- CarrierPoint) error {
	for d.More() { // iterate over points array
		point := CarrierPoint{Carrier: *carrier}
		_ = d.Decode(&point)
		ch <- point
	}
	return nil
}

func processCarriers(d *json.Decoder, ch chan<- CarrierPoint) error {
	for d.More() { // iterate over carrier array
		token, _ := d.Token()
		// check if token is start of object
		if token == json.Delim('{') {
			var carrier Carrier
			// read the object
			for d.More() {
				carrierToken, _ := d.Token()
				switch carrierToken {
				case "id":
					_ = d.Decode(&(carrier.ID))
				case "name":
					_ = d.Decode(&(carrier.Name))
				case "points":
					_, _ = d.Token() // array start
					_ = processCarrierPoints(d, &carrier, ch)
					_, _ = d.Token() // array end
				default:
					_ = skip(d)
				}
			}
			// closing object
			_, _ = d.Token()
		}
	}
	return nil
}

func ProcessFile(ch chan<- CarrierPoint) error {
	r, err := os.Open(filePath)
	if err != nil {
		return err
	}
	d := json.NewDecoder(r)
	// find carriers array
	for d.More() {
		token, err := d.Token()
		if err != nil {
			return err
		}
		if token == "carriers" {
			err = processCarriers(d, ch)
		}

	}
	return nil
}
