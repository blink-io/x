package orm

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/blink-io/opt/omit"
	"github.com/blink-io/opt/omitnull"
	"github.com/blink-io/sqx"
	"github.com/blink-io/x/id"
	"github.com/blink-io/x/ptr"
	"github.com/blink-io/x/sql/misc"
	"github.com/blink-io/x/types/tuplen"
	"github.com/bokwoon95/sq"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/spf13/cast"
	"github.com/stretchr/testify/require"
	"github.com/uptrace/bun"
)

func TestSq_Pg_Tag_Insert_1(t *testing.T) {
	db := getPgDBForSQ()
	tbl := Tables.Tags

	s1 := randomTag(nil).Setter()
	s2 := randomTag(Ptr(gofakeit.School())).Setter()
	rt, err := tbl.Insert(ctx, sq.Log(db), s1, s2)

	require.NoError(t, err)
	require.NotNil(t, rt)
}

func TestSq_Pg_Tag_Insert_2(t *testing.T) {
	db := getPgDBForSQ()
	tbl := Tables.Tags

	setter := randomTag(nil).Setter()

	_, err := tbl.Insert(ctx, sq.Log(db), setter)
	require.NoError(t, err)
}

func TestSq_Pg_Tag_Insert_3(t *testing.T) {
	db := getPgDBForSQ()
	tbl := Tables.Tags

	var iii = time.Now().Unix()
	s1 := randomTag(Ptr("")).Setter()
	s1.ID = omit.From(iii + 10)

	s2 := randomTag(Ptr(gofakeit.AppName())).Setter()
	s2.ID = omit.From(iii + 30)

	_, err := tbl.Insert(ctx, sq.Log(db), s1, s2)
	require.NoError(t, err)
}

func TestSq_Pg_Tag_Update_1(t *testing.T) {
	db := getPgDBForSQ()
	tbl := Tables.Tags

	var vid int64 = 40

	s := TagSetter{
		ID:   omit.From[int64](vid),
		Name: omit.From(gofakeit.City()),
		Code: omit.From(id.ShortID()),
	}

	where := sq.And(sqx.AlwaysTrueExpr, tbl.ID.EqInt64(vid))
	rt, err := tbl.Update(ctx, sq.Log(db), where, s)

	require.NoError(t, err)
	require.NotNil(t, rt)
}

func TestSq_Pg_Tag_Update_WithTx_1(t *testing.T) {
	db := getPgDBForSQ()
	tbl := Tables.Tags

	where := sq.And(sqx.AlwaysTrueExpr, tbl.ID.EqInt(15))

	runInTxFn := func(ctx context.Context, db sq.DB) error {
		row, err := tbl.One(ctx, sq.Log(db), where)
		if err != nil {
			return err
		}

		s := row.Setter()
		s.Name = omit.From(gofakeit.City() + "-Modified")
		s.Description = omitnull.From(gofakeit.Animal())

		rt, err := tbl.Update(ctx, sq.Log(db), where, s)
		if err != nil {
			return err
		}
		require.NotNil(t, rt)

		return nil
	}

	t.Run("tx with commit", func(t *testing.T) {
		err := sqx.RunInTx(ctx, db, nil, runInTxFn)
		require.NoError(t, err)
	})

	t.Run("tx with rollback", func(t *testing.T) {
		err := sqx.RunInTx(ctx, db, nil, func(ctx context.Context, tx sq.DB) error {
			_ = runInTxFn(ctx, tx)
			return errors.New("tx with rollback")
		})

		require.Error(t, err)
	})

	t.Run("tx with bun and commit", func(t *testing.T) {
		bundb := getPgDBForBun()
		require.NoError(t, bundb.HealthCheck(ctx))

		err := bundb.RunInTx(ctx, nil, func(ctx context.Context, tx bun.Tx) error {
			return runInTxFn(ctx, tx.Tx)
		})
		require.NoError(t, err)
	})
}

func TestSq_Pg_Tag_Update_WithBun_WithTx_1(t *testing.T) {
	t.Parallel()

	db := getPgDBForBun()
	tbl := Tables.Tags
	where := sq.And(sqx.AlwaysTrueExpr, tbl.ID.EqInt(15))

	runInTxFn := func(ctx context.Context, tx bun.Tx) error {
		row, err := tbl.One(ctx, sq.Log(tx.Tx), where)
		if err != nil {
			return err
		}

		s := row.Setter()
		s.Name = omit.From(gofakeit.City() + "-Modified")
		s.Description = omitnull.From(gofakeit.Animal())

		rt, err := tbl.Update(ctx, sq.Log(tx.Tx), where, s)
		if err != nil {
			return err
		}
		require.NotNil(t, rt)

		return nil
	}

	t.Run("tx with bun and commit", func(t *testing.T) {
		err := db.RunInTx(ctx, nil, func(ctx context.Context, tx bun.Tx) error {
			return runInTxFn(ctx, tx)
		})
		require.NoError(t, err)
	})

	t.Run("tx with bun and rollback", func(t *testing.T) {
		err := db.RunInTx(ctx, nil, func(ctx context.Context, tx bun.Tx) error {
			_ = runInTxFn(ctx, tx)
			return errors.New("tx with rollback")
		})

		require.Error(t, err)
	})
}

func TestSq_Pg_Tag_Update_WithSqlx_WithTx_1(t *testing.T) {
	t.Parallel()

	db := getPgDBForSqlx()
	tbl := Tables.Tags
	where := sq.And(sqx.AlwaysTrueExpr, tbl.ID.EqInt(15))

	db.Begin()

	runInTxFn := func(ctx context.Context, db sq.DB) error {
		row, err := tbl.One(ctx, sq.Log(db), where)
		if err != nil {
			return err
		}

		s := row.Setter()
		s.Name = omit.From(gofakeit.City() + "-Modified")
		s.Description = omitnull.From(gofakeit.Animal())

		rt, err := tbl.Update(ctx, sq.Log(db), where, s)
		if err != nil {
			return err
		}
		require.NotNil(t, rt)

		return nil
	}

	t.Run("tx with sqlx and commit", func(t *testing.T) {
		err := sqx.RunInTx(ctx, db, nil, runInTxFn)
		require.NoError(t, err)
	})

	t.Run("tx with sqlx and rollback", func(t *testing.T) {
		err := sqx.RunInTx(ctx, db, nil, func(ctx context.Context, db sq.DB) error {
			_ = runInTxFn(ctx, db)
			return errors.New("tx with rollback")
		})
		require.Error(t, err)
	})
}

func TestSq_Pg_Tag_One_1(t *testing.T) {
	db := getPgDBForSQ()
	tbl := Tables.Tags

	where := sq.And(sqx.AlwaysTrueExpr, tbl.ID.GtInt(0))

	row, err := tbl.One(ctx, sq.Log(db), where)

	require.NoError(t, err)
	require.NotNil(t, row)
}

func TestSq_Pg_Tag_One_2(t *testing.T) {
	db := getPgDBForSQ()
	tbl := Tables.Tags

	where := sq.And(sqx.AlwaysTrueExpr, tbl.ID.GtInt(1000))

	fields := []sq.Field{
		tbl.ID,
		tbl.NAME,
		tbl.CODE,
	}

	q := sq.From(tbl).Select(fields...).Where(where)
	row, err := sq.FetchOne(sq.Log(db), q, func(r *sq.Row) map[string]string {
		m := map[string]string{
			"id":   cast.ToString(r.Int64(tbl.ID.GetName())),
			"name": r.String(tbl.NAME.GetName()),
			"code": r.String(tbl.CODE.GetName()),
		}
		return m
	})

	require.NoError(t, err)
	require.NotNil(t, row)
}

func TestSq_Pg_Tag_All_1(t *testing.T) {
	db := getPgDBForSQ()
	tbl := Tables.Tags

	where := sq.And(sqx.AlwaysTrueExpr, tbl.ID.GtInt(0))

	rows, err := tbl.All(ctx, sq.Log(db), where)

	require.NoError(t, err)
	require.NotNil(t, rows)
}

func TestSq_Pg_Tag_Delete_1(t *testing.T) {
	db := getPgDBForSQ()
	tbl := Tables.Tags

	where := sq.And(sqx.AlwaysTrueExpr, tbl.ID.GtInt(1000000))

	rows, err := tbl.Delete(ctx, sq.Log(db), where)

	require.NoError(t, err)
	require.NotNil(t, rows)
}

func TestSq_Pg_Tag_Executor_1(t *testing.T) {
	db := getPgDBForSQ()
	tbl := Tables.Tags
	exec := tbl.Executor()
	wh := tbl.ID.EqInt(40)
	r, err := exec.One(ctx, sq.Log(db), wh)
	require.NoError(t, err)
	require.NotNil(t, r)
}

func TestSq_Pg_Tag_Mapper_Insert_OnConflict_1(t *testing.T) {
	db := getPgDBForSQ()
	mm := Mappers.TAGS
	tbl := mm.Table()

	r1 := randomTag(nil)
	r1.ID = 1
	r2 := randomTag(Ptr(gofakeit.City()))
	r2.ID = 2
	r3 := randomTag(nil)
	r3.ID = 3

	nrs := []TagSetter{r1.Setter(), r2.Setter(), r3.Setter()}

	q := sq.Postgres.InsertInto(tbl).
		ColumnValues(mm.InsertT(ctx, nrs...)).
		OnConflict(tbl.ID).
		DoUpdateSet(tbl.DESCRIPTION.SetString("DoUpdateSet"))

	rt, err := sq.ExecContext(ctx, sq.Log(db), q)
	require.NoError(t, err)

	fmt.Println(rt)
}

func TestSq_Pg_Tag_Mapper_Update_1(t *testing.T) {
	db := getPgDBForSQ()
	mm := Mappers.TAGS
	tbl := mm.Table()

	s := TagSetter{
		ID:   omit.From[int64](15),
		Name: omit.From(gofakeit.City()),
		Code: omit.From(id.ShortID()),
	}

	q := sq.Update(tbl).
		SetFunc(mm.UpdateT(ctx, s)).
		Where(sqx.AlwaysTrueExpr, sqx.AlwaysTrueExpr, tbl.ID.EqInt(15))

	rt, err := sq.Exec(sq.Log(db), q)
	require.NoError(t, err)
	require.NotNil(t, rt)
}

func TestSq_Pg_Tag_Mapper_Insert_Returning_1(t *testing.T) {
	db := getPgDBForSQ()
	tbl := Tables.Tags
	mm := Mappers.TAGS

	nrs := []TagSetter{
		randomTag(nil).Setter(),
		randomTag(Ptr(gofakeit.City())).Setter(),
	}

	rr, err := sq.FetchAll(db, sq.Postgres.
		InsertInto(tbl).
		ColumnValues(mm.InsertT(ctx, nrs...)),
		mm.SelectT(ctx),
	)

	require.NoError(t, err)
	require.NotNil(t, rr)
}

func TestSq_Pg_Tag_Insert_Select_1(t *testing.T) {
	db := getPgDBForSQ()
	tbl := Tables.TagsBak
	fTbl := Tables.Tags

	_, err := sq.Exec(db, sq.Postgres.
		InsertInto(tbl).
		Columns(tbl.GUID, tbl.NAME, tbl.CODE, tbl.DESCRIPTION, tbl.CREATED_AT).
		Select(sq.
			Select(fTbl.GUID, fTbl.NAME, fTbl.CODE, fTbl.DESCRIPTION, fTbl.CREATED_AT).
			From(fTbl).
			Where(fTbl.DESCRIPTION.IsNotNull()),
		).
		SetDialect(sq.DialectPostgres),
	)
	require.NoError(t, err)
}

func TestSq_Pg_Tag_Mapper_Insert_1(t *testing.T) {
	db := getPgDBForSQ()
	mm := Mappers.TAGS
	tbl := mm.Table()

	d1 := randomTag(nil).Setter()
	d2 := randomTag(ptr.Of("Hello, Hi, 你好")).Setter()

	_, err := sq.Exec(sq.Log(db), sq.
		InsertInto(tbl).
		ColumnValues(mm.InsertT(ctx, d1, d2)),
	)
	require.NoError(t, err)
}

func TestSq_Pg_Tag_Mapper_FetchAll_1(t *testing.T) {
	db := getPgDBForSQ()
	mm := Mappers.TAGS
	tbl := mm.Table()

	query := sq.
		From(tbl).
		Where(tbl.ID.GtInt(0)).
		Limit(100)
	records, err := sq.FetchAllContext(ctx, db, query, mm.SelectT(ctx))

	require.NoError(t, err)
	require.NotNil(t, records)
}

func TestSq_Pg_Tag_Fetch_Custom_1(t *testing.T) {
	db := getPgDBForSQ()
	mm := Mappers.TAGS
	tbl := mm.Table()

	query := sq.
		Queryf("select id, code, name, {} as status from tags limit 5", sq.Literal("OK"))
	records, err := sq.FetchAll(db, query, func(r *sq.Row) tuplen.Tuple4[int, string, string, string] {
		return tuplen.Of4(
			r.Int(tbl.ID.GetName()),
			r.String(tbl.CODE.GetName()),
			r.String(tbl.NAME.GetName()),
			r.String("status"),
		)
	})

	require.NoError(t, err)
	require.NotNil(t, records)
}

func TestSq_Pg_Tag_Mapper_FetchAll_Paging(t *testing.T) {
	db := getPgDBForSQ()
	mm := Mappers.TAGS
	tbl := mm.Table()
	perPage := 3
	qm := mm.SelectT(ctx)
	vdb := sq.VerboseLog(db)

	bq := sq.
		From(tbl).
		Where(tbl.ID.GtInt(0)).
		Limit(perPage).
		OrderBy(tbl.ID.Asc())

	t.Run("pagination", func(t *testing.T) {
		q1 := bq.Offset(misc.ToOffset(1, perPage))
		rs1, err1 := sq.FetchAll(vdb, q1, qm)
		require.NoError(t, err1)
		require.NotNil(t, rs1)

		q2 := bq.Offset(misc.ToOffset(2, perPage))
		rs2, err2 := sq.FetchAll(vdb, q2, qm)
		require.NoError(t, err2)
		require.NotNil(t, rs2)

		q3 := bq.Offset(misc.ToOffset(3, perPage))
		rs3, err3 := sq.FetchAll(vdb, q3, qm)
		require.NoError(t, err3)
		require.NotNil(t, rs3)
	})

	fmt.Println("done")
}

func TestSq_Pg_Tag_Mapper_FetchOne_ByPK(t *testing.T) {
	db := getPgDBForSQ()
	mm := Mappers.TAGS
	tbl := mm.Table()

	idWhere := tbl.PrimaryKeyValues(23)
	query := sq.
		From(tbl).
		Where(idWhere)
	records, err := sq.FetchOne(db, query, mm.SelectT(ctx))

	require.NoError(t, err)
	require.NotNil(t, records)
}
