package offset

type Pagination struct {
	Offset  int  `json:"offset"`
	Limit   int  `json:"limit"`
	HasMore bool `json:"has_more"`
}

type Result[E any] struct {
	Pagination Pagination `json:"pagination"`
	Records    []E        `json:"records"`
}
