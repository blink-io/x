package page

var (
	defaultPage    = 1
	defaultPerPage = 10
)

type Pagination struct {
	Page         int  `json:"page"`
	PerPage      int  `json:"per_page"`
	TotalPages   int  `json:"total_pages,omitempty"`
	TotalRecords int  `json:"total_records,omitempty"`
	HasMore      bool `json:"has_more"`
}

type Result[E any] struct {
	Pagination Pagination `json:"pagination"`
	Records    []E        `json:"records"`
}

// NewResult creates a page-base result.
// If page is 0, use 1 as default page; if perPage is 0, use 10 as default per_page.
func NewResult[E any](page int, perPage int, totalRecords int, records []E) Result[E] {
	if page == 0 {
		page = defaultPage
	}
	if perPage == 0 {
		perPage = defaultPerPage
	}
	p := Pagination{
		Page:    page,
		PerPage: perPage,
	}
	if totalRecords > 0 {
		var totalPages int
		n := totalRecords % perPage
		if n > 0 {
			totalPages = int((totalRecords-n)/perPage) + 1
		} else {
			totalPages = (totalRecords - n) / perPage
		}
		p.TotalPages = totalPages
		p.TotalRecords = totalRecords
		p.HasMore = page < totalPages
	} else {
		// If records are empty, no more data
		p.HasMore = len(records) > 0
	}

	m := Result[E]{
		Pagination: p,
		Records:    records,
	}
	return m
}
