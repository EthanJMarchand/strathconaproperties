package sp

type ListingFilter struct {
	HighPrice     string
	LowPrice      string
	HighBeds      string
	lowBeds       string
	HighBaths     string
	LowBaths      string
	HighYearBuilt string
	LowYearBuilt  string
}

type Listing struct {
	ID        string `json:"id"`
	Bedrooms  string `json:"bedrooms"`
	Bathrooms string `json:"bathrooms"`
}

type ListingService interface {
	Filter(filter *ListingFilter) ([]Listing, error)
}
