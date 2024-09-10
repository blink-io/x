package orm

import "time"

type Array struct {
	ID         int64          `db:"id"`
	StrArrays  []string       `db:"str_arrays"`
	Int4Arrays []int32        `db:"int4_arrays"`
	BoolArrays []bool         `db:"bool_arrays"`
	CreatedAt  time.Time      `db:"created_at"`
	VJsonb     map[string]any `db:"v_jsonb"`
	VJson      map[string]any `db:"v_json"`
}
