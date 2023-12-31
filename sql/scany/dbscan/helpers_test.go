package dbscan_test

import (
	"database/sql"
	"reflect"
	"testing"

	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/blink-io/x/sql/scany/dbscan"
)

type testModel struct {
	Foo string
	Bar string
}

const (
	multipleRowsQuery = `
		SELECT *
		FROM (
			VALUES ('foo val', 'bar val'), ('foo val 2', 'bar val 2'), ('foo val 3', 'bar val 3')
		) AS t (foo, bar)
	`
	singleRowsQuery = `
		SELECT 'foo val' AS foo, 'bar val' AS bar
	`
)

// RowsAdapter makes pgx.Rows compliant with the dbscan.Rows interface.
// See dbscan.Rows for details.
type RowsAdapter struct {
	pgx.Rows
}

// NewRowsAdapter returns a new RowsAdapter instance.
func NewRowsAdapter(rows pgx.Rows) *RowsAdapter {
	return &RowsAdapter{Rows: rows}
}

// Columns implements the dbscan.Rows.Columns method.
func (ra RowsAdapter) Columns() ([]string, error) {
	columns := make([]string, len(ra.Rows.FieldDescriptions()))
	for i, fd := range ra.Rows.FieldDescriptions() {
		columns[i] = fd.Name
	}
	return columns, nil
}

// Close implements the dbscan.Rows.Close method.
func (ra RowsAdapter) Close() error {
	ra.Rows.Close()
	return nil
}

func makeStrPtr(v string) *string { return &v }

func queryRows(t *testing.T, query string) dbscan.Rows {
	t.Helper()
	pgxRows, err := testDB.Query(ctx, query)
	require.NoError(t, err)
	rows := NewRowsAdapter(pgxRows)
	return rows
}

func getAPI(opts ...dbscan.APIOption) (*dbscan.API, error) {
	if len(opts) < 1 {
		opts = []dbscan.APIOption{}
	}
	opts = append(opts, dbscan.WithScannableTypes(
		(*sql.Scanner)(nil),
	))
	return dbscan.NewAPI(opts...)
}

func scan(t *testing.T, dst interface{}, rows dbscan.Rows) error {
	defer rows.Close() //nolint: errcheck
	rs := testAPI.NewRowScanner(rows)
	rows.Next()
	if err := rs.Scan(dst); err != nil {
		return err
	}
	requireNoRowsErrorsAndClose(t, rows)
	return nil
}

func requireNoRowsErrorsAndClose(t *testing.T, rows dbscan.Rows) {
	t.Helper()
	require.NoError(t, rows.Err())
	require.NoError(t, rows.Close())
}

func allocateDestination(v interface{}) interface{} {
	dstType := reflect.TypeOf(v)
	dst := reflect.New(dstType).Interface()
	return dst
}

func assertDestinationEqual(t *testing.T, expected, dst interface{}) {
	t.Helper()
	got := reflect.ValueOf(dst).Elem().Interface()
	assert.Equal(t, expected, got)
}
