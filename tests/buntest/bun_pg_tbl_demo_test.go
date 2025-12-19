package buntest

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/sanity-io/litter"

	"github.com/blink-io/x/misc/closer"
	"github.com/stretchr/testify/require"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
)

func TestPg_TblDemo(t *testing.T) {
	ctx := context.Background()
	db := getPgDB()
	bundb := bun.NewDB(db, pgdialect.New(),
		bun.WithDiscardUnknownColumns(),
	)
	defer closer.CloseQuietly(bundb.Close)

	t.Run("insert multi", func(t *testing.T) {
		rr := []*TblDemoSetter{
			randomTblDemoSetter(),
			randomTblDemoSetter(),
		}

		q := bundb.NewInsert().Model(&rr).Returning("id")
		_, err := q.Exec(ctx)
		require.NoError(t, err)
	})

	t.Run("insert multi 2", func(t *testing.T) {
		rr := []*TblDemo{
			randomTblDemo(),
			randomTblDemo(),
		}

		q := bundb.NewInsert().Model(&rr).Returning("id")
		_, err := q.Exec(ctx)
		require.NoError(t, err)
	})

	t.Run("select rows", func(t *testing.T) {
		var rr []*TblDemo
		count, err := bundb.NewSelect().
			Model((*TblDemo)(nil)).
			Order("id asc").
			//Where("n_json is not null").
			ScanAndCount(ctx, &rr)
		require.NoError(t, err)

		require.Equal(t, len(rr), count)
		fmt.Println(count)
	})

	t.Run("select custom", func(t *testing.T) {
		type Info struct {
			V string `json:"v"`
			N int32  `json:"n"`
		}
		var rr []json.RawMessage
		err := bundb.NewSelect().
			Column("v_json").
			Table("tbl_demo").
			//Where("n_json is not null").
			Where("id = ?", 1).
			Scan(ctx, &rr)
		for idx, r := range rr {
			_ = idx
			var i Info
			_ = json.Unmarshal(r, &i)
			fmt.Println(litter.Sdump(i))
		}
		require.NoError(t, err)
	})
}
