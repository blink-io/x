package orm

import (
	"context"
	"testing"
	"time"

	"github.com/blink-io/opt/omit"
	"github.com/blink-io/sq"
	"github.com/stretchr/testify/require"
)

func TestPg_Enums_Insert_1(t *testing.T) {
	db := GetPgDB()

	d1 := EnumSetter{
		Status:    omit.From(EnumEnumsStatusActive),
		CreatedAt: omit.From(time.Now()),
	}
	d2 := EnumSetter{
		Status:    omit.From(EnumEnumsStatusBlocked),
		CreatedAt: omit.From(time.Now()),
	}

	var exec = Executors.Enum
	_, err := exec.Insert(ctx, db, []EnumSetter{d1, d2}...)
	require.NoError(t, err)
}

func Test_Enums_SelectAll_1(t *testing.T) {
	db := GetPgDB()
	var exec = Executors.Enum
	where := exec.Table().ID.GtInt(0)
	results, err := exec.All(ctx, db, where)
	require.NoError(t, err)
	require.True(t, len(results) > 0)
}

func Test_Enums_SelectCustom_1(t *testing.T) {
	db := GetPgDB()
	var exec = Executors.Enum
	tbl := exec.Table()
	where := exec.Table().ID.GtInt(0)

	q := sq.Select(tbl.ID, tbl.STATUS).From(tbl).Where(where)
	rr, err := sq.FetchAll[EnumEnumsStatus](db, q, func(ctx context.Context, row *sq.Row) EnumEnumsStatus {
		var status EnumEnumsStatus
		row.Enum(&status, tbl.STATUS.GetName())
		return status
	})
	require.NoError(t, err)
	require.True(t, len(rr) > 0)
}
