package packeta

import "strconv"

type Carrier struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type CarrierPoint struct {
	Code    string `json:"code"`
	Street  string `json:"street"`
	City    string `json:"city"`
	Carrier Carrier
}

func (cp *CarrierPoint) String() string {
	return cp.Code + " " + cp.Street + " " + cp.City + " " + cp.Carrier.Name + " " + strconv.Itoa(cp.Carrier.ID)
}
