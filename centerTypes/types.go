package centerTypes

type CowinCenters struct {
	Centers []Center `json:"centers"`
}

type Center struct {
	Name         string        `json:"name"`
	Address      string        `json:"address"`
	Block_name   string        `json:"block_name"`
	Pincode      int           `json:"pincode"`
	Lat          int           `json:"lat"`
	Long         int           `json:"long"`
	From         string        `json:"from"`
	To           string        `json:"to"`
	Fee_type     string        `json:"fee_type"`
	Sessions     []Session     `json:"sessions"`
	Vaccine_fees []Vaccine_fee `json:"vaccine_fees"`
}

type Session struct {
	Available_capacity       int      `json:"available_capacity"`
	Slots                    []string `json:"slots"`
	Min_age_limit            int      `json:"min_age_limit"`
	Vaccine                  string   `json:"vaccine"`
	Available_capacity_dose1 int      `json:"available_capacity_dose1"`
	Available_capacity_dose2 int      `json:"available_capacity_dose2"`
}

type Vaccine_fee struct {
	Vaccine string `json:"vaccine"`
	Fee     string `json:"fee"`
}
