package orm

import (
	"testing"

	"github.com/blink-io/x/ptr"
	"github.com/bokwoon95/sq"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/stretchr/testify/require"
)

func TestSq_Mysql_Tag_Mapper_Insert_1(t *testing.T) {
	db := getMysqlDBForSQ()
	mm := Mappers.TAGS
	tbl := Tables.Tags

	d1 := randomTag(nil).Setter()
	d2 := randomTag(ptr.Of("Hello, Hi, 你好")).Setter()

	q := sq.InsertInto(tbl).
		ColumnValues(mm.InsertT(ctx, d1, d2))

	_, err := sq.Exec(sq.VerboseLog(db), q)
	require.NoError(t, err)
}

func TestSq_Mysql_Tag_Update_1(t *testing.T) {
	db := getMysqlDBForSQ()
	mm := Mappers.TAGS
	tbl := mm.Table()

	idWhere := tbl.PrimaryKeyValues(4)
	q := sq.Update(tbl).Where(idWhere).
		SetFunc(func(c *sq.Column) {
			c.SetString(tbl.DESCRIPTION, gofakeit.DomainName())
		}).
		Set(tbl.CODE.SetString(gofakeit.UUID()))

	_, err := sq.Exec(sq.VerboseLog(db), q)
	require.NoError(t, err)
}

func TestSq_Mysql_Tag_Update_2(t *testing.T) {
	db := getMysqlDBForSQ()
	mm := Mappers.TAGS
	tbl := mm.Table()

	idWhere := tbl.PrimaryKeyValues(4)
	q := sq.Update(tbl).
		Where(idWhere).
		Set(tbl.CODE.SetString(gofakeit.UUID()))

	_, err := sq.Exec(sq.VerboseLog(db), q)
	require.NoError(t, err)
}

func TestSq_Mysql_Tag_FetchOne_ByID(t *testing.T) {
	db := getMysqlDBForSQ()
	mm := Mappers.TAGS
	tbl := mm.Table()

	idWhere := tbl.PrimaryKeyValues(4)
	query := sq.
		From(tbl).Where(idWhere)

	records, err := sq.FetchOne(sq.VerboseLog(db), query, func(r *sq.Row) Tag {
		return Tag{
			ID:   r.Int64("id"),
			Code: r.String("code"),
			Name: r.String("name"),
			GUID: r.String("guid"),
			//CreatedAt: r.TimeField(tbl.CREATED_AT),
		}
	})

	require.NoError(t, err)
	require.NotNil(t, records)
}

func TestSq_Mysql_Tag_FetchAll_BySQL_1(t *testing.T) {
	db := getMysqlDBForSQ()
	idb := sq.VerboseLog(db)
	//tbl := TagTable

	query := sq.Queryf("select t.* from tags t where t.id = {} and t.code = {}", 4, "Philadelphia")

	records, err := sq.FetchOne(idb, query, func(r *sq.Row) Tag {
		return Tag{
			ID:   r.Int64("id"),
			Code: r.String("code"),
			Name: r.String("name"),
			GUID: r.String("guid"),
			//CreatedAt: r.TimeField(tbl.CREATED_AT),
		}
	})

	require.NoError(t, err)
	require.NotNil(t, records)
}

func TestSq_Mysql_FetchAll_BySQL_1(t *testing.T) {
	db := getMysqlDBForSQ()
	idb := sq.VerboseLog(db)
	//tbl := TagTable

	query := sq.Queryf("SELECT '{{}' as code, '{{abcd}' as name")

	records, err := sq.FetchOne(idb, query, func(r *sq.Row) Tag {
		return Tag{
			//UserID:   r.Int("c1"),
			Code: r.String("code"),
			Name: r.String("name"),
			//UserGUID: r.String("guid"),
			//CreatedAt: r.TimeField(tbl.CREATED_AT),
		}
	})

	require.NoError(t, err)
	require.NotNil(t, records)
}
