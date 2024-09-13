package token

type Pagination[T any] struct {
	NextToken T    `json:"next_token"`
	Size      int  `json:"size"`
	HasMore   bool `json:"has_more"`
}

type Result[T any, E any] struct {
	Pagination Pagination[T] `json:"pagination"`
	Records    []E           `json:"records"`
}
