package orm

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log/slog"
	"testing"
	"time"

	"github.com/blink-io/opt/omit"
	"github.com/bokwoon95/sq"
	"github.com/brianvoe/gofakeit/v7"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/stretchr/testify/require"
)

func TestSq_1(t *testing.T) {
	db := GetPgDB()

	qr := sq.Postgres.Queryf("select 'heison' as name, version() as version")
	m, err := sq.FetchOne[*Model](db, qr, func(row *sq.Row) *Model {
		return &Model{
			Name:    row.String("name"),
			Version: row.String("version"),
		}
	})
	require.NoError(t, err)
	require.NotNil(t, m)

	fmt.Println(m)
}

func TestSq_Pg_Insert_User_1(t *testing.T) {
	db := GetPgDB()

	now := time.Now()
	_, err := sq.Exec(db, sq.
		InsertInto(UserTable).
		Columns(
			UserTable.GUID,
			UserTable.USERNAME,
			UserTable.SCORE,
			UserTable.CREATED_AT,
			UserTable.UPDATED_AT,
		).
		Values(
			gofakeit.UUID(),
			gofakeit.Username(),
			gofakeit.Float64(),
			now,
			now,
		).
		Values(
			gofakeit.UUID(),
			gofakeit.Username(),
			gofakeit.Float64(),
			now,
			now,
		).
		Values(
			gofakeit.UUID(),
			gofakeit.Username(),
			gofakeit.Float64(),
			now,
			now,
		),
	)
	require.NoError(t, err)

	gofakeit.Date()
}

func TestSq_Pg_User_Insert_1(t *testing.T) {
	db := GetPgDB()
	tbl := Tables.Users

	records := []User{
		randomUser(),
		randomUser(),
		randomUser(),
	}

	_, err := sq.Exec(db,
		sq.InsertInto(tbl).
			ColumnValues(tbl.InsertMapper(records...)),
	)
	require.NoError(t, err)
}

func TestSq_Pg_User_FetchAll_WithTenantID_1(t *testing.T) {
	db := GetPgDB()
	tbl := Tables.Users
	vctx := context.WithValue(ctx, tbl.TENANT_ID.GetName(), 3)

	query := sq.Postgres.From(tbl).Where(tbl.ID.GtInt(0)).Limit(100)
	records, err := sq.FetchAllContext(vctx, db, query, tbl.QueryMapper())

	require.NoError(t, err)
	require.NotNil(t, records)
}

func TestSq_Pg_User_FetchAll_2(t *testing.T) {
	db := GetPgDB()
	tbl := Tables.Users

	query := sq.Postgres.From(tbl).Where(tbl.ID.GtInt(0)).Limit(100)
	records, err := sq.FetchAllContext(ctx, sq.Log(db), query, tbl.QueryMapper())

	require.NoError(t, err)
	require.NotNil(t, records)
}

func TestSq_Pg_UserDevice_Insert_ColumnMapper_1(t *testing.T) {
	db := GetPgDB()
	tbl := UserDeviceTable

	records := []*UserDevice{
		randomUserDevice(),
		randomUserDevice(),
		randomUserDevice(),
	}

	_, err := sq.Exec(db, sq.
		InsertInto(tbl).ColumnValues(func(col *sq.Column) {
		for _, r := range records {
			userDeviceInsertColumnMapper(col, r)
		}
	}))
	require.NoError(t, err)
}

func TestSq_Pg_UserWithDevice_FetchAll_Join_1(t *testing.T) {
	db := GetPgDB()
	tbl := UserTable
	joinTbl := UserDeviceTable
	sb := sq.Postgres

	fields, rowMapper := userWithDeviceSelect()
	query := sb.
		Select(
			fields...,
		).
		From(tbl).
		LeftJoin(joinTbl, tbl.ID.Eq(joinTbl.USER_ID)).
		Where(tbl.ID.GtInt(0)).
		Limit(100).
		OrderBy(tbl.GUID.Desc())
	rs, err := sq.FetchAllContext(ctx, db, query, rowMapper)

	require.NoError(t, err)
	require.NotNil(t, rs)
}

func TestSq_Pg_User_Update_1(t *testing.T) {
	db := GetPgDB()
	tbl := UserTable

	_, err := sq.Exec(db, sq.
		Update(tbl).
		SetFunc(func(col *sq.Column) {
			col.SetString(tbl.USERNAME, "DAN")
			col.SetFloat64(tbl.SCORE, gofakeit.Float64Range(50, 80))
		}).
		Where(tbl.ID.EqInt64(2)),
	)
	require.NoError(t, err)
}

func TestSq_Pg_User_Update_2(t *testing.T) {
	db := GetPgDB()

	var us UserSetter
	us.ID = omit.From[int](10)
	us.Score = omit.From(gofakeit.Float64Range(55, 90))
	us.Username = omit.From[string](gofakeit.Username() + "-Modified")

	require.NoError(t, us.Update(db))
}

func TestSq_Pg_User_Delete_1(t *testing.T) {
	db := GetPgDB()
	tbl := UserTable

	_, err := sq.Exec(db, sq.
		DeleteFrom(tbl).
		Where(tbl.ID.EqInt64(56)),
	)
	require.NoError(t, err)
}

func TestSq_Pg_Enum_Insert_1(t *testing.T) {
	db := GetPgDB()
	tbl := sq.New[ENUMS]("e")
	_, err := sq.Exec(db, sq.
		InsertInto(tbl).
		Columns(tbl.STATUS).
		Values(UserStatusBlocked).
		Values(UserStatusActive),
	)
	require.NoError(t, err)
}

func TestSq_Pg_Enum_Insert_Tx_Success_1(t *testing.T) {
	db := GetPgDB()
	tbl := sq.New[ENUMS]("e")

	tx, err := db.Begin()
	require.NoError(t, err)

	defer handleTxPanic(tx)

	_, err = sq.Exec(sq.Log(tx), sq.
		InsertInto(tbl).
		Columns(tbl.STATUS).
		Values(UserStatusActive),
	)

	_, err = sq.Exec(sq.Log(tx), sq.
		InsertInto(tbl).
		Columns(tbl.STATUS).
		Values(UserStatusBlocked),
	)

	require.NoError(t, err)
}

func TestSq_ShowAll(t *testing.T) {
	db := GetPgDB()

	sql := "show all"
	qr := sq.Postgres.Queryf(sql)
	m, err := sq.FetchAll(db, qr, func(row *sq.Row) map[string]string {
		return map[string]string{
			"name":        row.String("name"),
			"setting":     row.String("setting"),
			"description": row.String("description"),
		}
	})
	require.NoError(t, err)
	require.NotNil(t, m)

	fmt.Println(m)
}

func handleTxPanic(tx *sql.Tx) {
	if r := recover(); r != nil {
		errx := tx.Rollback()
		if errx != nil {
			slog.Info("Rollback err")
		}
		slog.Error("do rollback for tx")
	} else {
		_ = tx.Commit()
		slog.Info("do commit for tx")
	}
}

func TestSq_Pg_Enum_Insert_Tx_Fail_1(t *testing.T) {
	db := GetPgDB()
	tbl := sq.New[ENUMS]("e")

	tx, err := db.Begin()
	require.NoError(t, err)

	defer handleTxPanic(tx)

	doPanic := gofakeit.RandomInt([]int{2, 4, 6})%2 == 0
	_, err = sq.Exec(sq.Log(tx), sq.
		InsertInto(tbl).
		Columns(tbl.STATUS).
		Values(UserStatusActive),
	)

	if doPanic {
		panic(errors.New("panic for tx"))
	}

	_, err = sq.Exec(sq.Log(tx), sq.
		InsertInto(tbl).
		Columns(tbl.STATUS).
		Values(UserStatusBlocked),
	)

	require.NoError(t, err)
}

func TestSq_Pg_UserWithDevice_Exists_1(t *testing.T) {
	db := GetPgDB()
	tbl := Tables.Users
	joinTbl := Tables.UserDevices
	sb := sq.Postgres

	query := sb.
		Select().
		From(tbl).
		Where(sq.Exists(sq.SelectOne().From(joinTbl).Where(joinTbl.USER_ID.Eq(tbl.ID)))).
		Limit(100).
		OrderBy(tbl.GUID.Desc())
	rs, err := sq.FetchAllContext(ctx, db, query, userModelRowMapper())

	require.NoError(t, err)
	require.NotNil(t, rs)
}

func TestSq_Pg_UserWithDevice_CTE_1(t *testing.T) {
	db := GetPgDB()
	tbl := Tables.Users
	joinTbl := Tables.UserDevices
	sb := sq.Postgres

	// create the CTE
	devicesCTE := sq.NewCTE("devices", nil, sq.Postgres.
		Select(
			joinTbl.USER_ID,
		).
		From(joinTbl).
		GroupBy(joinTbl.USER_ID),
	)

	query := sb.
		With(devicesCTE).
		From(tbl).
		Join(devicesCTE, devicesCTE.Field("user_id").Eq(tbl.ID)).
		Limit(100).
		OrderBy(tbl.GUID.Desc())
	rs, err := sq.FetchAllContext(ctx, db, query, userModelRowMapper())

	require.NoError(t, err)
	require.NotNil(t, rs)
}

func Ptr[T any](v T) *T {
	return &v
}
