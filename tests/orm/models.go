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

type User struct {
	ID        int       `db:"id"`
	GUID      string    `db:"guid"`
	Username  string    `db:"username"`
	FirstName string    `db:"first_name"`
	LastName  string    `db:"last_name"`
	Score     float64   `db:"score"`
	Level     int       `db:"level"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	TenantID  int       `db:"tenant_id"`
}
