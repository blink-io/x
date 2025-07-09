package buntest

import (
	"context"
	"testing"

	"github.com/blink-io/x/misc/closer"
	"github.com/stretchr/testify/require"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
)

func TestPg_TbleDemo(t *testing.T) {
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
}
