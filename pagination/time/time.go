package time

type Pagination struct {
	Before  int64 `json:"before"`
	After   int64 `json:"after"`
	Size    int   `json:"size"`
	HasMore bool  `json:"has_more"`
}

type Result[E any] struct {
	Pagination Pagination `json:"pagination"`
	Records    []E        `json:"records"`
}
