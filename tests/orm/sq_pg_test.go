package orm

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"log/slog"
	"sync"
	"testing"
	"time"

	"github.com/aarondl/opt/omit"
	sqx "github.com/blink-io/x/sql/orm/sq"
	"github.com/bokwoon95/sq"
	"github.com/brianvoe/gofakeit/v7"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/stretchr/testify/require"
)

var pgOnce sync.Once

func getPgDBForSQ() *sql.DB {
	return getPgDB()
}

func getPgDB() *sql.DB {
	pgOnce.Do(func() {
		setupPgDialect()
	})

	dsn := "postgres://test:test@192.168.50.88:5432/test?sslmode=disable"
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		log.Fatalf("failed to open pg db: %v", err)
	}
	if err := db.Ping(); err != nil {
		log.Fatalf("failed to ping pg db: %v", err)
	}

	return db
}

func setupPgDialect() {
	dialect := sq.DialectPostgres
	sqx.SetDefaultDialect(dialect)
	slog.Info("Setup database dialect", "dialect", dialect)
}

func TestSq_1(t *testing.T) {
	db := getPgDBForSQ()

	qr := sq.SQLite.Queryf("select 'heison' as name, version() as version")
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
	db := getPgDBForSQ()

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
	db := getPgDBForSQ()
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
	db := getPgDBForSQ()
	tbl := Tables.Users
	vctx := context.WithValue(ctx, tbl.TENANT_ID.GetName(), 3)

	query := sq.Postgres.From(tbl).Where(tbl.ID.GtInt(0)).Limit(100)
	records, err := sq.FetchAllContext(vctx, db, query, tbl.QueryMapper())

	require.NoError(t, err)
	require.NotNil(t, records)
}

func TestSq_Pg_User_FetchAll_2(t *testing.T) {
	db := getPgDBForSQ()
	tbl := Tables.Users

	query := sq.Postgres.From(tbl).Where(tbl.ID.GtInt(0)).Limit(100)
	records, err := sq.FetchAllContext(ctx, sq.Log(db), query, tbl.QueryMapper())

	require.NoError(t, err)
	require.NotNil(t, records)
}

func TestSq_Pg_UserDevice_Insert_ColumnMapper_1(t *testing.T) {
	db := getPgDBForSQ()
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
	db := getPgDBForSQ()
	tbl := UserTable
	joinTbl := UserDeviceTable

	fields, rowMapper := userWithDeviceSelect()
	query := sq.
		Select(
			fields...,
		).
		From(tbl).
		Join(joinTbl, tbl.ID.Eq(joinTbl.USER_ID)).
		Where(tbl.ID.GtInt(0)).
		Limit(100)
	records, err := sq.FetchAll(db, query, rowMapper)

	require.NoError(t, err)
	require.NotNil(t, records)
}

func TestSq_Pg_User_Update_1(t *testing.T) {
	db := getPgDBForSQ()
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
	db := getPgDBForSQ()

	var us UserSetter
	us.ID = omit.From[int](10)
	us.Score = omit.From(gofakeit.Float64Range(55, 90))
	us.Username = omit.From[string](gofakeit.Username() + "-Modified")

	require.NoError(t, us.Update(db))
}

func TestSq_Pg_User_Delete_1(t *testing.T) {
	db := getPgDBForSQ()
	tbl := UserTable

	_, err := sq.Exec(db, sq.
		DeleteFrom(tbl).
		Where(tbl.ID.EqInt64(56)),
	)
	require.NoError(t, err)
}

func TestSq_Pg_Enum_Insert_1(t *testing.T) {
	db := getPgDBForSQ()
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
	db := getPgDBForSQ()
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
	db := getPgDBForSQ()
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

func Ptr[T any](v T) *T {
	return &v
}
