package orm

import (
	"time"

	"github.com/gofrs/uuid/v5"
	guuid "github.com/google/uuid"
)

type Array struct {
	ID          int64          `db:"id"`
	StrArrays   []string       `db:"str_arrays"`
	Int4Arrays  []int32        `db:"int4_arrays"`
	BoolArrays  []bool         `db:"bool_arrays"`
	CreatedAt   time.Time      `db:"created_at"`
	VJsonb      map[string]any `db:"v_jsonb"`
	VJson       map[string]any `db:"v_json"`
	VUUID       uuid.UUID      `db:"v_uuid"`
	VUUID2      guuid.UUID     `db:"v_uuid"`
	JsonArrays  []string       `db:"json_arrays"`
	JsonbArrays []string       `db:"jsonb_arrays"`
	UUIDArrays  []string       `db:"uuid_arrays"`
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
