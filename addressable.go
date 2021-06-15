package intercom

// AddressableList is an object used to contain multiple Addressable.
type AddressableList struct {
	// The type of object - list.
	Type string `json:"type"`
	// An array of Addressable Objects. Maximum of 10.
	Data []Addressable `json:"data"`
	// The URL where the full list can be accessed (ie. /contacts/1234/companies).
	Url string `json:"url"`
	// The total amount of records.
	TotalCount int64 `json:"total_count"`
	// Whether there's more Addressable Objects to be viewed. If true, use the url to view all.
	HasMore bool `json:"has_more"`
}

// Addressable is an object used to link a company, note or tag to a contact.
type Addressable struct {
	// The type of object - company, note, tag.
	Type string `json:"type"`
	// The id of the object.
	ID string `json:"id"`
	// The URL where the object in question can be accessed (ie. /companies/45678).
	Url string `json:"url"`
}
