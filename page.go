package intercom

// PageParams determine paging information to and from the API
type PageParams struct {
	Page          int64  `json:"page" url:"page,omitempty"`
	StartingAfter string `json:"starting_after" url:"starting_after,omitempty"`
	Next          *struct {
		Page          int64  `json:"page"`
		StartingAfter string `json:"starting_after"`
	} `json:"next" url:"-"`
	PerPage    int64 `json:"per_page" url:"per_page,omitempty"`
	TotalPages int64 `json:"total_pages" url:"-"`
}
