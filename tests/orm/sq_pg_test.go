package orm

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	"github.com/aarondl/opt/omit"
	"github.com/bokwoon95/sq"
	"github.com/brianvoe/gofakeit/v7"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/stretchr/testify/require"
)

func getPgDB() *sql.DB {
	dialect := sq.DialectPostgres
	sq.DefaultDialect.Store(&dialect)

	dsn := "postgres://blink:888asdf%21%23%25@192.168.50.88:5432/orm-demo?sslmode=disable"
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		log.Fatalf("failed to open db: %v", err)
	}
	if err := db.Ping(); err != nil {
		log.Fatalf("failed to ping db: %v", err)
	}
	return db
}

func init() {
	l := sq.NewLogger(os.Stdout, "", log.LstdFlags, sq.LoggerConfig{
		ShowTimeTaken: true,
		HideArgs:      true,
	})
	sq.SetDefaultLogQuery(func(ctx context.Context, stats sq.QueryStats) {
		l.SqLogQuery(ctx, stats)
	})
	dialect := sq.DialectPostgres
	sq.DefaultDialect.Store(&dialect)
}

func TestSq_1(t *testing.T) {
	db := getPgDB()

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
	db := getPgDB()

	now := time.Now()
	_, err := sq.Exec(db, sq.
		InsertInto(UserTableDef).
		Columns(
			//UserTableDef.ID,
			UserTableDef.GUID,
			UserTableDef.Username,
			UserTableDef.Score,
			UserTableDef.CreatedAt,
			UserTableDef.UpdatedAt,
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

func TestSq_Pg_Insert_UserDevice_1(t *testing.T) {
	db := getPgDB()
	tbl := UserDeviceTableDef
	now := time.Now()

	_, err := sq.Exec(db, sq.
		InsertInto(tbl).
		Columns(
			tbl.UserID,
			tbl.GUID,
			tbl.Device,
			tbl.Model,
			tbl.CreatedAt,
			tbl.UpdatedAt,
		).
		Values(
			gofakeit.IntRange(1, 30),
			gofakeit.UUID(),
			gofakeit.AppName(),
			gofakeit.CarModel(),
			now,
			now,
		).
		Values(
			gofakeit.IntRange(1, 30),
			gofakeit.UUID(),
			gofakeit.AppName(),
			gofakeit.CarModel(),
			now,
			now,
		).
		Values(
			gofakeit.IntRange(1, 30),
			gofakeit.UUID(),
			gofakeit.AppName(),
			gofakeit.CarModel(),
			now,
			now,
		),
	)
	require.NoError(t, err)
}

func TestSq_Pg_User_Insert_ColumnMapper_1(t *testing.T) {
	db := getPgDB()
	tbl := UserTableDef

	records := []*User{
		randomUser(),
		randomUser(),
		randomUser(),
	}

	_, err := sq.Exec(db, sq.
		InsertInto(tbl).ColumnValues(func(col *sq.Column) {
		for _, r := range records {
			userInsertColumnMapper(col, r)
		}
	}))
	require.NoError(t, err)
}

func TestSq_Pg_UserDevice_Insert_ColumnMapper_1(t *testing.T) {
	db := getPgDB()
	tbl := UserDeviceTableDef

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

func TestSq_Pg_User_FetchAll_1(t *testing.T) {
	db := getPgDB()
	tbl := UserTableDef

	query := sq.Postgres.From(tbl).Where(tbl.ID.GtInt(0)).Limit(100)
	records, err := sq.FetchAll(db, query, userModelRowMapper())

	require.NoError(t, err)
	require.NotNil(t, records)
}

func TestSq_Pg_User_Update_1(t *testing.T) {
	db := getPgDB()
	tbl := UserTableDef

	_, err := sq.Exec(db, sq.
		Update(tbl).
		SetFunc(func(col *sq.Column) {
			col.SetString(tbl.Username, "DAN")
			col.SetFloat64(tbl.Score, 0.88)
		}).
		Where(tbl.ID.EqInt64(2)),
	)
	require.NoError(t, err)
}

func TestSq_Pg_User_Update_2(t *testing.T) {
	db := getPgDB()

	var us UserSetter
	us.ID = omit.From[int](10)
	us.Score = omit.From(gofakeit.Float64Range(55, 90))
	us.Username = omit.From[string](gofakeit.Username() + "-Modified")

	require.NoError(t, us.Update(db))
}

func TestSq_Pg_User_Update_3(t *testing.T) {
	db := getPgDB()

	var us UserSetter
	//us.ID = omit.From[int](10)
	us.Score = omit.From(gofakeit.Float64Range(88, 90))
	//us.Username = omit.From[string](gofakeit.Username() + "-Modified")

	require.NoError(t, us.UpdateByWhere(db, UserTableDef.ID.LtFloat64(50)))
}

func TestSq_Pg_User_Delete_1(t *testing.T) {
	db := getPgDB()
	tbl := UserTableDef

	_, err := sq.Exec(db, sq.
		DeleteFrom(tbl).
		Where(tbl.ID.EqInt64(56)),
	)
	require.NoError(t, err)
}

func TestSq_Pg_Tag_Insert_1(t *testing.T) {
	db := getPgDB()

	tbl := TagTableDef

	records := []*Tag{
		randomTag(nil),
		randomTag(Ptr(gofakeit.City())),
		randomTag(nil),
	}

	_, err := sq.Exec(db, sq.
		InsertInto(tbl).ColumnValues(func(col *sq.Column) {
		for _, r := range records {
			tagInsertColumnMapper(col, r)
		}
	}))

	require.NoError(t, err)
}

func TestSq_Pg_Tag_Insert_2(t *testing.T) {
	db := getPgDB()

	err := randomTag(nil).Insert(db)
	require.NoError(t, err)

	err = randomTag(Ptr(gofakeit.School())).Insert(db)
	require.NoError(t, err)
}

func TestSq_Pg_Enum_Insert_1(t *testing.T) {
	db := getPgDB()
	tbl := sq.New[ENUMS]("e")
	_, err := sq.Exec(db, sq.
		InsertInto(tbl).
		Columns(tbl.STATUS).
		Values("blacocked"),
	)
	require.NoError(t, err)
}

func Ptr[T any](v T) *T {
	return &v
}
