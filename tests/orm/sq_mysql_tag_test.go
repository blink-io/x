package orm

import (
	"testing"

	"github.com/blink-io/x/ptr"
	"github.com/bokwoon95/sq"
	"github.com/stretchr/testify/require"
)

func TestSq_Mysql_Tag_Mapper_Insert_1(t *testing.T) {
	db := getMysqlDBForSQ()
	mm := NewTagMapper()
	tbl := Tables.Tags

	d1 := randomTag(nil)
	d2 := randomTag(ptr.Of("Hello, Hi, 你好"))

	_, err := sq.Exec(sq.VerboseLog(db), sq.
		InsertInto(tbl).
		ColumnValues(mm.InsertMapper(d1, d2)),
	)
	require.NoError(t, err)
}

func TestSq_Mysql_Tag_FetchOne_ByID(t *testing.T) {
	db := getMysqlDBForSQ()
	idb := sq.VerboseLog(db)
	tbl := TagTable

	idWhere := tbl.PrimaryKeyValues(4)
	query := sq.
		From(tbl).Where(idWhere)

	records, err := sq.FetchOne(idb, query, func(r *sq.Row) Tag {
		return Tag{
			ID:   r.Int("id"),
			Code: r.String("code"),
			Name: r.String("name"),
			GUID: r.String("guid"),
			//CreatedAt: r.TimeField(tbl.CREATED_AT),
		}
	})

	require.NoError(t, err)
	require.NotNil(t, records)
}
