package orm

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/url"
	"os"
	"testing"
	"time"

	"github.com/bokwoon95/sq"
	"github.com/brianvoe/gofakeit/v7"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/sanity-io/litter"
	"github.com/stretchr/testify/require"
	"github.com/uptrace/bun"
)

type UserTable struct {
	sq.TableStruct `sq:"users"`
	ID             sq.NumberField `sq:"id"`
	GUID           sq.StringField `sq:"guid"`
	Username       sq.StringField `sq:"username"`
	Score          sq.NumberField `sq:"score"`
	CreatedAt      sq.TimeField   `sq:"created_at"`
	UpdatedAt      sq.TimeField   `sq:"updated_at"`
}

type UserDeviceTable struct {
	sq.TableStruct `sq:"user_devices"`
	ID             sq.NumberField `sq:"id"`
	GUID           sq.StringField `sq:"guid"`
	UserID         sq.NumberField `sq:"user_id"`
	Device         sq.StringField `sq:"device"`
	Model          sq.StringField `sq:"model"`
	CreatedAt      sq.TimeField   `sq:"created_at"`
	UpdatedAt      sq.TimeField   `sq:"updated_at"`
}

type User struct {
	ID        int       `db:"id"`
	GUID      string    `db:"guid"`
	Username  string    `db:"username"`
	Score     float64   `db:"score"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func (m User) String() string {
	return litter.Sdump(m)
}

type UserDevice struct {
	ID        int       `db:"id"`
	GUID      string    `db:"guid"`
	UserID    int       `db:"user_id"`
	Device    string    `db:"device"`
	Model     string    `db:"score"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func (m UserDevice) String() string {
	return litter.Sdump(m)
}

type Model struct {
	Name    string       `db:"name"`
	Version string       `db:"version"`
	Now     bun.NullTime `db:"now"`
}

func (m Model) String() string {
	return litter.Sdump(m)
}

var _ fmt.Stringer = (*User)(nil)
var _ fmt.Stringer = (*UserDevice)(nil)

var userTable = sq.New[UserTable]("u1")
var userDeviceTable = sq.New[UserDeviceTable]("u2")

func getPgDB() *sql.DB {
	dsn := "postgres://blink:888asdf%21%23%25@192.168.50.88:5432/orm_demo?sslmode=disable"
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
		InsertInto(userTable).
		Columns(
			//userTable.ID,
			userTable.GUID,
			userTable.Username,
			userTable.Score,
			userTable.CreatedAt,
			userTable.UpdatedAt,
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
	tbl := userDeviceTable
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
	tbl := userTable

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
	tbl := userDeviceTable

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
	tbl := userTable

	query := sq.Postgres.From(tbl).Where(tbl.ID.GtInt(0)).Limit(100)
	records, err := sq.FetchAll(db, query, userModelRowMapper())

	require.NoError(t, err)
	require.NotNil(t, records)
}

func TestSq_Pg_User_Update_1(t *testing.T) {
	db := getPgDB()
	tbl := userTable

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

func TestSq_Pg_User_Delete_1(t *testing.T) {
	db := getPgDB()
	tbl := userTable

	_, err := sq.Exec(db, sq.
		DeleteFrom(tbl).
		Where(tbl.ID.EqInt64(56)),
	)
	require.NoError(t, err)
}

func userModelRowMapper() func(*sq.Row) *User {
	return func(r *sq.Row) *User {
		tbl := userTable

		u := &User{
			ID:       r.IntField(tbl.ID),
			GUID:     r.StringField(tbl.GUID),
			Username: r.StringField(tbl.Username),
			Score:    r.Float64Field(tbl.Score),
		}

		dd := sq.DefaultDialect.Load()
		dz := sq.DialectSQLite
		if dz == *dd {
			crstr := r.String("created_at")
			upstr := r.String("updated_at")
			if ct, err := time.Parse(time.RFC3339Nano, crstr); err == nil {
				u.CreatedAt = ct
			}
			if ct, err := time.Parse(time.RFC3339Nano, upstr); err == nil {
				u.UpdatedAt = ct
			}
		} else {
			u.CreatedAt = r.TimeField(tbl.CreatedAt)
			u.UpdatedAt = r.TimeField(tbl.UpdatedAt)
		}

		return u
	}
}

func userInsertColumnMapper(col *sq.Column, r *User) {
	tbl := userTable

	col.Set(tbl.GUID, r.GUID)
	col.Set(tbl.Username, r.Username)
	col.Set(tbl.Score, r.Score)
	col.Set(tbl.CreatedAt, r.CreatedAt)
	col.Set(tbl.UpdatedAt, r.UpdatedAt)
}

func userDeviceInsertColumnMapper(col *sq.Column, r *UserDevice) {
	tbl := userDeviceTable

	col.Set(tbl.UserID, r.UserID)
	col.Set(tbl.GUID, r.GUID)
	col.Set(tbl.Device, r.Device)
	col.Set(tbl.Model, r.Model)
	col.SetTime(tbl.CreatedAt, r.CreatedAt)
	col.SetTime(tbl.UpdatedAt, r.UpdatedAt)
}

func TestURL_1(t *testing.T) {
	u := &url.URL{
		Scheme:   "postgres",
		Host:     "192.168.50.88:5432",
		Path:     "orm_demo",
		User:     url.UserPassword("blink", "888asdf!#%"),
		RawQuery: "sslmode=disable",
	}

	fmt.Println(u)
	fmt.Println(u.RequestURI())
}

func randomUser() *User {
	ln := time.Now().Local()
	u := &User{
		GUID:      gofakeit.UUID(),
		Username:  gofakeit.Username(),
		Score:     gofakeit.Float64(),
		CreatedAt: ln,
		UpdatedAt: ln,
	}
	return u
}

func randomUserDevice() *UserDevice {
	ln := time.Now().Local()
	u := &UserDevice{
		UserID:    gofakeit.IntRange(1, 30),
		GUID:      gofakeit.UUID(),
		Device:    gofakeit.AppName(),
		Model:     gofakeit.CarModel(),
		CreatedAt: ln,
		UpdatedAt: ln,
	}
	return u
}
