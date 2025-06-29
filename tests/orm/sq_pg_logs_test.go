package orm

import (
	"github.com/blink-io/opt/omit"
	"github.com/blink-io/opt/omitnull"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestPg_Logs_Insert_1(t *testing.T) {
	db := GetPgDB()

	d1 := LogSetter{
		ID:        omit.From(int64(gofakeit.Int32())),
		Content:   omit.From(gofakeit.BuzzWord()),
		CreatedAt: omit.From(time.Now()),
		AuditedAt: omitnull.From(time.Now()),
	}
	_, err := Executors.Log.Insert(ctx, db, []LogSetter{d1}...)
	require.NoError(t, err)
}
