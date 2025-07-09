package buntest

import (
	"testing"
	"time"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"golang.org/x/net/context"

	sq "github.com/Masterminds/squirrel"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/stretchr/testify/require"
)

func TestPg_Squirrel_Insert_1(t *testing.T) {
	db := getPgDB()
	cols := []string{
		TblSimpleTable.Columns.GUID,
		TblSimpleTable.Columns.Name,
		TblSimpleTable.Columns.CreatedAt,
	}
	vals := []any{
		gofakeit.UUID(),
		gofakeit.Name(),
		time.Now(),
	}
	_, err := sq.Insert(TblSimpleTable.Name).
		Columns(cols...).
		Values(vals...).
		RunWith(db).
		PlaceholderFormat(sq.Dollar).
		Exec()
	require.Nil(t, err)
}

func TestPg_Users_SelectJoin_1(t *testing.T) {
	db := getPgDB()
	bundb := bun.NewDB(db, pgdialect.New())
	ctx := context.Background()

	type UserWithDevices struct {
		UserID        int64  `bun:"user_id"`
		UserFirstName string `bun:"user_first_name"`
		UserLastName  string `bun:"user_last_name"`
		UserDeviceID  int64  `bun:"user_device_id"`
		DeviceModel   string `bun:"device_model"`
		DeviceName    string `bun:"device_name"`
	}

	t.Run("users join user_devices", func(t *testing.T) {
		var records []UserWithDevices
		q := bundb.NewSelect().
			ColumnExpr("u.id as user_id").
			ColumnExpr("u.first_name as user_first_name").
			ColumnExpr("u.last_name as user_last_name").
			ColumnExpr("d.id as user_device_id").
			ColumnExpr("d.model as device_model").
			ColumnExpr("d.name as device_name").
			TableExpr("users u").
			Join("join user_devices d on u.id = d.user_id").
			Order("u.id asc")
		err := q.Scan(ctx, &records)
		require.NoError(t, err)
	})

}
