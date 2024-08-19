package ksql

import (
	"context"
	"fmt"
	"log/slog"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/vingarcia/ksql"
	"github.com/vingarcia/ksql/adapters/modernc-ksqlite"
)

var ctx = context.Background()

func init() {
	ksql.InjectLogger(ctx, func(ctx context.Context, values ksql.LogValues) {
		slog.InfoContext(ctx, "SQL",
			slog.String("query", values.Query),
		)
	})
}

type Info struct {
	Version string `json:"version" ksql:"version"`
}

type User struct {
	ID       int64  `ksql:"id"`
	Username string `ksql:"username"`
	GUID     string `ksql:"guid"`
	Location string `ksql:"location"`
}

func getSqliteDB(t *testing.T) ksql.DB {
	url := "./sqlite_demo.db"

	db, err := ksqlite.New(ctx, url, ksql.Config{MaxOpenConns: 1})
	require.NoError(t, err)
	require.NotNil(t, db)

	return db
}

func TestKsql_Insert_1(t *testing.T) {
	db := getSqliteDB(t)

	var info Info

	err := db.QueryOne(ctx, &info, "select sqlite_version() as version")
	require.NoError(t, err)

	fmt.Println(info.Version)
}

func TestKsql_Select_1(t *testing.T) {
	db := getSqliteDB(t)

	var users []*User

	err := db.Query(ctx, &users, "select * from users limit ?", 5)
	require.NoError(t, err)

	for _, user := range users {
		fmt.Printf("id: %d, username: %s, guid: %s\n", user.ID, user.Username, user.GUID)
	}
}
