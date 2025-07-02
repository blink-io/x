package orm

import (
	"context"
	"fmt"
	"testing"

	"github.com/blink-io/sq"
	"github.com/stretchr/testify/require"
)

type TS_ARRAYS struct {
	sq.TableStruct
	ID  sq.NumberField `ddl:"type=bigint notnull primarykey default=nextval('arrays_id_seq'::regclass)"`
	TSA sq.ArrayField  `ddl:"type=timestampz[]"`
}

func TestPg_TSZ_Insert_1(t *testing.T) {
	db := GetPgDB()
	var tsa = sq.New[TS_ARRAYS]("")

	q := sq.Queryf("select tsa from ts_arrays limit 1").SetDialect(sq.DialectPostgres)
	rr, err := sq.FetchOne[[]string](sq.VerboseLog(db), q, func(ctx context.Context, row *sq.Row) []string {
		fmt.Printf("%#v\n", row.Columns()) // []string{"actor_id", "first_name", "lname"}
		fmt.Printf("%#v\n", row.Values())  // []any{18, "DAN", "TORN"}
		var tt []string
		row.Array(&tt, tsa.TSA.GetName())
		return tt
	})
	require.NoError(t, err)
	fmt.Println(rr)
}
