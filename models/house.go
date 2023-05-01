package models

type CreateHouseReq struct {
	HouseName     string   `json:"house_name"`
	Address       string   `json:"address"`
	City          string   `json:"city"`
	State         string   `json:"state"`
	Country       string   `json:"country"`
	Pincode       string   `json:"pincode"`
	Description   string   `json:"description"`
	Bedrooms      int      `json:"bedrooms"`
	Bathrooms     int      `json:"bathrooms"`
	Price         int      `json:"price"`
	SquareFootage int      `json:"square_footage"`
	Amenities     []string `json:"amenities"`
	Pictures      []string `json:"pictures"`
}

type HouseResponse struct {
	HousePid      string   `json:"house_pid"`
	HouseName     string   `json:"house_name"`
	Address       string   `json:"address"`
	City          string   `json:"city"`
	State         string   `json:"state"`
	Country       string   `json:"country"`
	Pincode       string   `json:"pincode"`
	Description   string   `json:"description"`
	Bedrooms      int      `json:"bedrooms"`
	Bathrooms     int      `json:"bathrooms"`
	Price         int      `json:"price"`
	SquareFootage int      `json:"square_footage"`
	Amenities     []string `json:"amenities"`
	Pictures      []string `json:"pictures"`
}

type HouseFilter struct {
	City          string
	State         string
	Country       string
	Pincode       string
	Bedrooms      int
	Bathrooms     int
	MinPrice      int
	MaxPrice      int
	SquareFootage int
}
