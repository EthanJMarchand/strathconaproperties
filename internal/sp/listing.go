package sp

type ListingFilter struct {
	HighPrice     int
	LowPrice      int
	HighBeds      int
	lowBeds       int
	HighBaths     int
	LowBaths      int
	HighYearBuilt int
	LowYearBuilt  int
}

type Listing struct {
	ID           int    `json:"id"`
	Price        int    `json:"price"`
	StreetNumber string `json:"street_number"`
	StreetName   string `json:"street_name"`
	Beds         int    `json:"bedrooms"`
	Baths        int    `json:"bathrooms"`
	YearBuilt    int    `json:"YearBuilt"`
	CourtesyOf   string `json:"courtesyOf"`
}

type ListingService interface {
	Filter(filter *ListingFilter) ([]Listing, error)
}
