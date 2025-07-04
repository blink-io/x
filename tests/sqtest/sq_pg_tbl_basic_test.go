package sqtest

import (
	"context"
	"database/sql"
	"fmt"
	"testing"

	"github.com/blink-io/sq"
	"github.com/blink-io/x/ptr"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

const (
	EnumTblBasicNEnumActive  EnumTblBasicNEnum = "active"
	EnumTblBasicNEnumBlocked EnumTblBasicNEnum = "blocked"

	EnumTblBasicVEnumActive  EnumTblBasicVEnum = "active"
	EnumTblBasicVEnumBlocked EnumTblBasicVEnum = "blocked"
)

func init() {
	EnumTblBasicNEnumValues = append(EnumTblBasicNEnumValues,
		string(EnumTblBasicNEnumActive),
		string(EnumTblBasicNEnumBlocked))

	EnumTblBasicVEnumValues = append(EnumTblBasicVEnumValues,
		string(EnumTblBasicVEnumActive),
		string(EnumTblBasicVEnumBlocked))
}

func randomTblBasicSetter() TblBasicSetter {

	vuuid := uuid.New()
	vstr := gofakeit.Name()
	vnum := EnumTblBasicVEnumActive
	vjson := map[string]any{"v": gofakeit.City()}
	var s = TblBasicSetter{
		NUUID:  ptr.Of(sq.ValidFrom[[16]byte](uuid.New())),
		NEnum:  ptr.Of(sql.Null[EnumTblBasicNEnum]{V: EnumTblBasicNEnumActive, Valid: true}),
		NStr:   ptr.Of(sq.NullString{}),
		NTime:  ptr.Of(sq.NullTime{}),
		NInt32: ptr.Of(sq.ValidFrom(gofakeit.Int32())),

		VUUID:  ptr.Of[[16]byte](vuuid),
		VStr:   &vstr,
		VJson:  &vjson,
		VInt32: ptr.Of(gofakeit.Int32()),
		VEnum:  ptr.Of(vnum),
		VTime:  ptr.Of(gofakeit.Date()),
	}

	return s
}

func TestSqPg_TblBasic_Insert_1(t *testing.T) {
	db := getPgDB()
	tbl := Tables.TblBasic

	var ss = []TblBasicSetter{
		randomTblBasicSetter(),
		randomTblBasicSetter(),
	}

	q := sq.InsertInto(tbl).ColumnValues(tbl.ColumnMapper(ss...))
	_, err := sq.Exec(sq.VerboseLog(db), q)
	require.NoError(t, err)
}

func TestSqPg_TblBasic_FetchAll_1(t *testing.T) {
	db := getPgDB()
	tbl := Tables.TblBasic

	q := sq.Select().From(tbl)

	rr, err := sq.FetchAll[TblBasic](sq.VerboseLog(db), q, tbl.RowMapper)
	require.NoError(t, err)
	require.True(t, len(rr) > 0)
}

func TestSqPg_TblBasic_FetchAll_Enum_1(t *testing.T) {
	db := getPgDB()
	tbl := Tables.TblBasic

	fields := sq.Fields{
		tbl.ID,
		tbl.V_ENUM,
		tbl.N_ENUM,
	}
	q := sq.Select(fields...).From(tbl).Where(tbl.ID.In([]int64{3, 4}))

	rr, err := sq.FetchAll[TblBasic](sq.VerboseLog(db), q, func(ctx context.Context, row *sq.Row) TblBasic {
		v := TblBasic{}
		return v
	})
	require.NoError(t, err)
	require.True(t, len(rr) > 0)
}

func TestSqPg_TblBasic_FetchAll_Bytes_1(t *testing.T) {
	db := getPgDB()
	tbl := Tables.TblBasic

	fields := sq.Fields{
		tbl.ID,
		tbl.V_BYTES,
		tbl.N_BYTES,
	}
	q := sq.Select(fields...).From(tbl).Where(tbl.ID.In([]int64{3, 4}))

	rr, err := sq.FetchAll[TblBasic](sq.VerboseLog(db), q, func(ctx context.Context, row *sq.Row) TblBasic {
		v := TblBasic{}

		nBytes := row.NullBytes(tbl.N_BYTES.GetName())
		vBytes := row.Bytes(tbl.V_BYTES.GetName())

		v.VBytes = vBytes
		v.NBytes = nBytes
		return v
	})
	require.NoError(t, err)
	require.True(t, len(rr) > 0)
}

func TestSqPg_TblBasic_FetchAll_Array_1(t *testing.T) {
	db := getPgDB()
	tbl := Tables.TblBasic

	fields := sq.Fields{
		tbl.ID,
		tbl.V_STR_ARRAYS,
		tbl.N_STR_ARRAYS,
	}
	q := sq.Select(fields...).From(tbl).Where(tbl.ID.In([]int64{3, 4}))

	rr, err := sq.FetchAll[TblBasic](sq.VerboseLog(db), q, func(ctx context.Context, row *sq.Row) TblBasic {
		v := TblBasic{}
		v.VStrArrays = sq.ArrayFrom[[]string](row, tbl.V_STR_ARRAYS.GetName())
		v.NStrArrays = sq.NullArrayFrom[[]string](row, tbl.N_STR_ARRAYS.GetName())

		return v
	})
	require.NoError(t, err)
	require.True(t, len(rr) > 0)
}

func TestSqPg_TblBasic_FetchAll_JSON_1(t *testing.T) {
	db := getPgDB()
	tbl := Tables.TblBasic

	fields := sq.Fields{
		tbl.ID,
		tbl.V_JSON,
		tbl.N_JSON,
	}
	q := sq.Select(fields...).From(tbl).Where(tbl.ID.In([]int64{3, 4}))

	rr, err := sq.FetchAll[TblBasic](sq.VerboseLog(db), q, func(ctx context.Context, row *sq.Row) TblBasic {
		v := TblBasic{}

		nJSON := row.NullJSON(tbl.N_JSON.GetName())
		vJSON := row.JSON(tbl.V_JSON.GetName())

		v.VJson = vJSON
		v.NJson = nJSON

		vJsonBytes := row.JSONBytes(tbl.V_JSON.GetName())
		nJsonBytes := row.NullJSONBytes(tbl.N_JSON.GetName())

		fmt.Println(vJsonBytes)
		fmt.Println(nJsonBytes)

		return v
	})
	require.NoError(t, err)
	require.True(t, len(rr) > 0)
}
