package response

type GetAdressCount struct {
	Address string `json:"address"`
	Count   int    `json:"count"`
}

type GetOccupationCount struct {
	Occupation string `json:"occupation"`
	Count      int    `json:"count"`
}
