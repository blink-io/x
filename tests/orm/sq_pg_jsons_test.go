package orm

import (
	"context"
	"fmt"
	"github.com/blink-io/opt/null"
	"github.com/blink-io/sq"
	"github.com/blink-io/x/ptr"
	"github.com/stretchr/testify/require"
	"testing"
)

type JSONS struct {
	sq.TableStruct
	ID     sq.NumberField `ddl:"type=bigint notnull primarykey default=nextval('enums_id_seq'::regclass)"`
	V_JSON sq.JSONField   `ddl:"type=json"`
	V_UUID sq.UUIDField   `ddl:"type=uuid"`
}

type JSONSR struct {
	ID    int64
	Vuuid null.Val[[16]byte]
}

func TestPg_Jsons_Insert_1(t *testing.T) {
	db := GetPgDB()
	tbl := sq.New[JSONS]("")

	q := sq.InsertInto(tbl).ColumnValues(func(ctx context.Context, col *sq.Column) {
		var vUuid [16]byte
		col.SetUUID(tbl.V_UUID, vUuid)
	})
	_, err := sq.Exec(sq.VerboseLog(db), q)

	require.NoError(t, err)
}

func Test_Jsons_SelectAll_1(t *testing.T) {
	var zeroUUID = *new([16]byte)
	db := GetPgDB()
	tbl := sq.New[JSONS]("")

	q := sq.Select(tbl.ID, tbl.V_UUID).From(tbl)

	rr, err := sq.FetchAll[JSONSR](sq.VerboseLog(db), q, func(ctx context.Context, row *sq.Row) JSONSR {
		var rt = JSONSR{}
		var uv [16]byte
		row.UUID(&uv, tbl.V_UUID.GetName())

		fmt.Printf("%p\n", ptr.Of(zeroUUID))
		fmt.Printf("%p\n", ptr.Of(uv))
		require.False(t, ptr.Of(zeroUUID) == ptr.Of(uv))

		if zeroUUID != uv {
			rt.Vuuid = null.From(uv)
		}
		rt.ID = row.Int64(tbl.ID.GetName())
		return rt
	})
	require.NoError(t, err)
	require.True(t, len(rr) > 0)
}
