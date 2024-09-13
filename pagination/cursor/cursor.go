package cursor

type Pagination[T any] struct {
	Current    T    `json:"current"`
	NextCursor T    `json:"next_cursor"`
	Size       int  `json:"size"`
	HasMore    bool `json:"has_more"`
}

type Result[T any, E any] struct {
	Pagination Pagination[T] `json:"pagination"`
	Records    []E           `json:"records"`
}
