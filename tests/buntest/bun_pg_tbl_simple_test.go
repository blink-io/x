package buntest

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"

	"github.com/blink-io/opt/omit"
	"github.com/blink-io/opt/omitnull"
	"github.com/blink-io/x/misc/closer"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/sanity-io/litter"
	"github.com/stephenafamo/scan"
	"github.com/stephenafamo/scan/stdscan"
	"github.com/stretchr/testify/require"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
)

func TestBun_TblSimple_Tests(t *testing.T) {
	ctx := context.Background()
	db := getPgDB()
	bundb := bun.NewDB(db, pgdialect.New(),
		bun.WithDiscardUnknownColumns(),
	)

	defer closer.CloseQuietly(bundb.Close)

	bundb.RegisterModel(TblSimpleTable.Model)
	setupBunHooks(bundb)

	type DBInfo struct {
		Version string `bun:"version"`
	}

	bundb.RegisterModel(TblSimpleTable.Model)

	t.Run("insert", func(t *testing.T) {
		now := time.Now()
		ss := []*TblSimpleSetter{
			randomTblSimpleSetter(nil),
			randomTblSimpleSetter(nil),
			randomTblSimpleSetter(nil),
			randomTblSimpleSetter(&now),
			randomTblSimpleSetter(&now),
		}
		q := bundb.NewInsert().Model(&ss)
		tb := q.GetTableName()
		mm := q.GetModel()
		fmt.Println(tb)
		fmt.Println(mm)
		_, err := q.Exec(ctx)
		require.NoError(t, err)
	})

	t.Run("insert 2", func(t *testing.T) {
		now := time.Now()
		ss := []*TblSimpleSetter2{
			randomTblSimpleSetter2(nil),
			randomTblSimpleSetter2(nil),
			randomTblSimpleSetter2(nil),
			randomTblSimpleSetter2(&now),
		}
		q := bundb.NewInsert().Model(&ss)
		tb := q.GetTableName()
		mm := q.GetModel()
		fmt.Println(tb)
		fmt.Println(mm)
		_, err := q.Exec(ctx)
		require.NoError(t, err)
	})

	t.Run("update selective", func(t *testing.T) {
		s := randomTblSimpleSetter(nil)
		q := bundb.NewUpdate().Model(s).Column("name").WherePK()
		_, err := q.Exec(ctx)
		require.NoError(t, err)
	})

	t.Run("update omit zero", func(t *testing.T) {
		s := randomTblSimpleSetter(nil)
		q := bundb.NewUpdate().Model(s).Column("name").WherePK()
		_, err := q.OmitZero().Exec(ctx)
		require.NoError(t, err)
	})

	t.Run("update by map", func(t *testing.T) {
		m := map[string]any{
			//string(TblSimpleTable.Columns.ID):   11,
			string(TblSimpleTable.Columns.Name): gofakeit.Animal(),
		}
		q := bundb.NewUpdate().Model(&m).
			Table(string(TblSimpleTable.Name)).
			Where("id = ?", 11)
		_, err := q.Exec(ctx)
		require.NoError(t, err)
	})

	t.Run("select raw", func(t *testing.T) {
		q := bundb.NewRaw("select id, name from tbl_simple where id = ?", 11)
		var mm []map[string]any
		err := q.Scan(ctx, &mm)
		require.NoError(t, err)
		fmt.Println(mm)
	})

	type CustomResult struct {
		ID        int                 `bun:"id"`
		Name      string              `bun:"name"`
		CreatedAt time.Time           `bun:"created_at"`
		DeletedAt sql.Null[time.Time] `bun:"deleted_at"`
	}

	t.Run("select custom columns", func(t *testing.T) {
		q := bundb.NewSelect().
			Column("id", "name", "created_at", "deleted_at").
			Table(TblSimpleTable.Name)
		var mm []CustomResult
		err := q.Scan(ctx, &mm)
		require.NoError(t, err)
	})

	t.Run("select all", func(t *testing.T) {
		var rr []*TblSimple
		err := bundb.NewSelect().
			Model(TblSimpleTable.Model).
			Order("id desc").
			Scan(ctx, &rr)
		require.NoError(t, err)

		fmt.Println(rr)
	})

	t.Run("select rows", func(t *testing.T) {
		count, err := bundb.NewSelect().Model(TblSimpleTable.Model).Count(ctx)
		require.NoError(t, err)

		fmt.Println(count)
	})

	t.Run("select count", func(t *testing.T) {
		rows, err := bundb.NewSelect().
			Column("id", "name", "deleted_at").
			Where("deleted_at is not null").
			Model((*TblSimple)(nil)).Rows(ctx)
		require.NoError(t, err)

		for rows.Next() {
			var id int64
			var name string
			var deletedAt sql.Null[time.Time]
			_ = rows.Scan(&id, &name, &deletedAt)
			fmt.Printf("id: %d, name: %s\n, deletedAt: %#v\n", id, name, deletedAt)
		}
	})

	t.Run("delete by id", func(t *testing.T) {
		q := bundb.NewDelete().
			Model(TblSimpleTable.Model).
			Where("id = ?", 11)
		rr, err := q.Exec(ctx)
		require.NoError(t, err)
		fmt.Println(litter.Sdump(rr))
	})

	t.Run("delete by id 2", func(t *testing.T) {
		q := bundb.NewDelete().
			Table(TblSimpleTable.Name).
			Where("id = ?", 11)
		rr, err := q.Exec(ctx)
		require.NoError(t, err)
		fmt.Println(litter.Sdump(rr))
	})

	t.Run("select by scan", func(t *testing.T) {
		ver, err := stdscan.One(ctx, db, scan.SingleColumnMapper[string], "select version();")
		require.NoError(t, err)

		fmt.Println(ver)
	})

	t.Run("update array 1", func(t *testing.T) {
		vals := []string{gofakeit.Animal(), gofakeit.City()}
		q := bundb.NewUpdate().
			Set("str_arrays = ?", pgdialect.Array(vals)).
			Table(TblSimpleTable.Name).
			Where("id = ?", 59)
		_, err := q.Exec(ctx)
		require.NoError(t, err)
	})

	t.Run("update array 1", func(t *testing.T) {
		r := TblSimpleSetter{
			ID:    omit.From[int64](60),
			NJSON: omitnull.From(map[string]any{"animal": gofakeit.Animal()}),
		}
		q := bundb.NewUpdate().
			Model(&r).
			OmitZero().
			WherePK()
		_, err := q.Exec(ctx)
		require.NoError(t, err)
	})

	t.Run("update by map", func(t *testing.T) {
		um := map[string]any{
			//"id":     61,
			"n_json": map[string]any{"name": gofakeit.Name(), "animal": gofakeit.Animal()},
		}
		q := bundb.NewUpdate().
			Model(&um).
			Table(TblSimpleTable.Name).
			Where("id = ?", 61)
		_, err := q.Exec(ctx)
		require.NoError(t, err)
	})
}
