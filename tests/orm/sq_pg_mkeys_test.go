package orm

import (
	"github.com/blink-io/sq"
	"github.com/brianvoe/gofakeit/v7"
	"testing"
	"time"

	"github.com/blink-io/opt/omit"
	"github.com/stretchr/testify/require"
)

func TestPg_MKeys_Insert_1(t *testing.T) {
	db := GetPgDB()

	d1 := MkeySetter{
		Id1:       omit.From(gofakeit.Int32()),
		Id2:       omit.From(gofakeit.Int32()),
		Name:      omit.From(gofakeit.Name()),
		GUID:      omit.From(gofakeit.UUID()),
		CreatedAt: omit.From(time.Now()),
	}
	d2 := MkeySetter{
		Id1:       omit.From(gofakeit.Int32()),
		Id2:       omit.From(gofakeit.Int32()),
		Name:      omit.From(gofakeit.Name()),
		GUID:      omit.From(gofakeit.UUID()),
		CreatedAt: omit.From(time.Now()),
	}

	var exec = Executors.MKey
	_, err := exec.Insert(ctx, db, []MkeySetter{d1, d2}...)
	require.NoError(t, err)
}

func Test_MKeys_SelectAll_1(t *testing.T) {
	db := sq.VerboseLog(GetPgDB())
	var exec = Executors.MKey
	where := sq.And(Tables.Mkeys.ID1.GtInt(0), Tables.Mkeys.ID2.GtInt(0))
	results, err := exec.All(ctx, db, where)
	require.NoError(t, err)
	require.True(t, len(results) > 0)
}
