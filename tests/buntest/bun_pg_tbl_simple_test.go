package buntest

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/blink-io/x/misc/closer"
	"github.com/blink-io/x/ptr"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/sanity-io/litter"
	"github.com/stephenafamo/scan"
	"github.com/stephenafamo/scan/stdscan"
	"github.com/stretchr/testify/require"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"testing"
	"time"
)

func TestBun_TblSimple_Tests(t *testing.T) {
	ctx := context.Background()
	db := getPgDB()
	bundb := bun.NewDB(db, pgdialect.New(),
		bun.WithDiscardUnknownColumns(),
	)
	defer closer.CloseQuietly(bundb.Close)

	type DBInfo struct {
		Version string `bun:"version"`
	}

	bundb.RegisterModel(TblSimpleTable.Model)

	var randomTblSimpleSetter = func() *TblSimpleSetter {
		r := &TblSimpleSetter{
			//ID:        ptr.Of(int64(11)),
			Name:      ptr.Of("test"),
			GUID:      ptr.Of(gofakeit.UUID()),
			CreatedAt: ptr.Of(time.Now()),
		}
		return r
	}

	t.Run("insert", func(t *testing.T) {
		ss := []*TblSimpleSetter{
			randomTblSimpleSetter(),
			randomTblSimpleSetter(),
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
		s := &TblSimpleSetter{
			ID:   ptr.Of(int64(11)),
			Name: ptr.Of(gofakeit.Name()),
		}
		q := bundb.NewUpdate().Model(s).Column("name").WherePK()
		_, err := q.Exec(ctx)
		require.NoError(t, err)
	})

	t.Run("update omit zero", func(t *testing.T) {
		s := &TblSimpleSetter{
			ID:        ptr.Of(int64(11)),
			Name:      ptr.Of(gofakeit.Name()),
			DeletedAt: ptr.Of(sql.Null[time.Time]{V: time.Now(), Valid: true}),
		}
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
		ID        int `bun:"id"`
		Name      string
		CreatedAt time.Time           `bun:"created_at"`
		DeletedAt sql.Null[time.Time] `bun:"deleted_at"`
	}

	t.Run("select custom columns", func(t *testing.T) {
		q := bundb.NewSelect().Column("id", "name", "created_at", "deleted_at").Table(string(TblSimpleTable.Name))
		var mm []CustomResult
		err := q.Scan(ctx, &mm)
		require.NoError(t, err)
	})

	t.Run("select rows", func(t *testing.T) {
		count, err := bundb.NewSelect().Model((*TblSimple)(nil)).Count(ctx)
		require.NoError(t, err)

		fmt.Println(count)
	})

	t.Run("select count", func(t *testing.T) {
		rows, err := bundb.NewSelect().Column("id", "name", "deleted_at").
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
		q := bundb.NewDelete().Model((*TblSimple)(nil)).Where("id = ?", 11)
		rr, err := q.Exec(ctx)
		require.NoError(t, err)
		fmt.Println(litter.Sdump(rr))
	})

	t.Run("delete by id 2", func(t *testing.T) {
		q := bundb.NewDelete().Table(string(TblSimpleTable.Name)).Where("id = ?", 11)
		rr, err := q.Exec(ctx)
		require.NoError(t, err)
		fmt.Println(litter.Sdump(rr))
	})

	t.Run("select by scan", func(t *testing.T) {
		ver, err := stdscan.One(ctx, db, scan.SingleColumnMapper[string], "select version();")
		require.NoError(t, err)

		fmt.Println(ver)
	})
}
