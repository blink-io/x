package pagination

type OffsetLimit[E any] struct {
	Offset  int  `json:"offset"`
	Limit   int  `json:"limit"`
	HasMore bool `json:"has_more"`

	Records []E `json:"records"`
}

type CursorBased[C any, E any] struct {
	Current    C    `json:"current"`
	NextCursor C    `json:"next_cursor"`
	Size       int  `json:"size"`
	HasMore    bool `json:"has_more"`

	Records []E `json:"records"`
}

type TokenBased[T any, E any] struct {
	NextToken T    `json:"next_token"`
	Size      int  `json:"size"`
	HasMore   bool `json:"has_more"`

	Records []E `json:"records"`
}

type TimeBased[E any] struct {
	Before  int64 `json:"before"`
	After   int64 `json:"after"`
	Size    int   `json:"size"`
	HasMore bool  `json:"has_more"`

	Records []E `json:"records"`
}
