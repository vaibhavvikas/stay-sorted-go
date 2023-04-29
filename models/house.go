package models

type CreateHouseReq struct {
	HouseName   string `json:"house_name"`
	City        string `json:"city"`
	State       string `json:"state"`
	Country     string `json:"country"`
	Pincode     string `json:"pincode"`
	Description string `json:"description"`
	Bedrooms    int    `json:"bedrooms"`
	Bathrooms   int    `json:"bathrooms"`
	Price       int    `json:"price"`
	Amenities   string `json:"amenities"`
	Pictures    []struct {
		URL string `json:"url"`
	} `json:"pictures"`
}
