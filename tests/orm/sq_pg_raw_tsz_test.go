package orm

import (
	"fmt"
	"github.com/blink-io/sq"
	rawsq "github.com/bokwoon95/sq"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

type RAW_TS_ARRAYS struct {
	rawsq.TableStruct `sq:"ts_arrays"`
	ID                rawsq.NumberField `ddl:"type=bigint notnull primarykey default=nextval('arrays_id_seq'::regclass)"`
	TSA               rawsq.ArrayField  `ddl:"type=timestampz[]"`
}

func TestPg_RawTSZ_Insert_1(t *testing.T) {
	db := GetPgDB()

	q := rawsq.Queryf("select {} from ts_arrays limit 1", sq.Fields{sq.Expr("id"), sq.Expr("tsa")}).SetDialect(rawsq.DialectPostgres)
	rr, err := rawsq.FetchOne[[]time.Time](rawsq.VerboseLog(db), q, func(row *rawsq.Row) []time.Time {
		var tt []time.Time
		//row.Array(&tt, "tsa")
		fmt.Printf("%#v\n", row.Int("id"))
		return tt
	})
	require.NoError(t, err)
	fmt.Println(rr)
}
