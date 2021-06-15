package intercom

// Location is an object identifying a Location.
type Location struct {
	// The type of object - location.
	Type string `json:"type"`
	// The country where the contact is.
	Country string `json:"country"`
	// A subdivision of the country which the contact is in (ie. state, province, county, territory, etc).
	Region string `json:"region"`
	// The city where the contact is.
	City string `json:"city"`
}
