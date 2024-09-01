package orm

import (
	"fmt"
	"log/slog"
	"strings"
	"testing"
	"time"

	dbrslog "github.com/blink-io/x/orm/dbr/logger/slog"
	"github.com/gocraft/dbr/v2"
	"github.com/guregu/null/v5"
	"github.com/stretchr/testify/require"
)

func quoteIdent(s, quote string) string {
	part := strings.SplitN(s, ".", 2)
	if len(part) == 2 {
		return quoteIdent(part[0], quote) + "." + quoteIdent(part[1], quote)
	}
	return quote + s + quote
}

type sqlite3 struct{}

func (d sqlite3) QuoteIdent(s string) string {
	return quoteIdent(s, `"`)
}

func (d sqlite3) EncodeString(s string) string {
	// https://www.sqlite.org/faq.html
	return `'` + strings.Replace(s, `'`, `''`, -1) + `'`
}

func (d sqlite3) EncodeBool(b bool) string {
	// https://www.sqlite.org/lang_expr.html
	if b {
		return "1"
	}
	return "0"
}

func (d sqlite3) EncodeTime(t time.Time) string {
	// https://www.sqlite.org/lang_datefunc.html
	return `'` + t.Local().Format(time.RFC3339Nano) + `'`
}

func (d sqlite3) EncodeBytes(b []byte) string {
	// https://www.sqlite.org/lang_expr.html
	return fmt.Sprintf(`X'%x'`, b)
}

func (d sqlite3) Placeholder(_ int) string {
	return "?"
}

func TestDBR_1(t *testing.T) {
	sl := dbrslog.New(slog.Default(), slog.LevelInfo)

	db := getSqliteDB()
	dd := sqlite3{}
	cc := &dbr.Connection{
		DB:            db,
		EventReceiver: sl,
		Dialect:       dd,
	}

	sess := cc.NewSession(nil)

	m := new(Model)

	var _ null.Time

	err := sess.SelectBySql("select datetime('now') as now, 'heison' as name, sqlite_version() as version").
		LoadOne(m)
	require.NoError(t, err)

	fmt.Println(m)
}

func TestDBR_User_Insert_1(t *testing.T) {
	sl := dbrslog.New(slog.Default(), slog.LevelInfo)

	db := getSqliteDB()
	dd := sqlite3{}
	cc := &dbr.Connection{
		DB:            db,
		EventReceiver: sl,
		Dialect:       dd,
	}

	sess := cc.NewSession(nil)

	m1 := randomUser()
	m2 := randomUser()

	_, err := sess.InsertInto("users").
		Columns([]string{"username", "guid", "score", "created_at", "updated_at"}...).
		Record(m1).
		Record(m2).
		Exec()
	require.NoError(t, err)

	fmt.Println("done")
}

func TestTime_1(t *testing.T) {
	tt := time.Now()
	bb, err := tt.MarshalText()
	kk := tt.String()
	require.NoError(t, err)
	fmt.Println(string(bb))
	fmt.Println(kk)
}
